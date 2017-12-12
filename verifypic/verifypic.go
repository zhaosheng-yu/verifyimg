package verifypic

import (
	"encoding/json"
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	width     = 320                 //图片宽度
	height    = 100                 //图片高度
	iconwidth = 22                  //icon边长
	fontFile  = "./source/ant4.ttf" //字体样式
	bgpicDir  = "./source/"         //背景库
	fontSize  = 24                  //字体大小
	fontDPI   = 72                  //屏幕每英寸的分辨率
	validnum  = 3                   //图片上需要选择的文字个数
	cacheDir  = "./cache/"          //cacheType为file时 缓存目录
	cacheType = "file"
)

type FontDot struct {
	Dx   int
	Dy   int
	Text string
}

//icon图片输出
func Hoverclick(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")

	img, _ := ioutil.ReadFile(bgpicDir + "hoverclick.png")

	fmt.Fprintf(w, string(img))
}

//获取保存的文字顺序
func GetData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sessid := r.Form.Get("sessid")
	if sessid == "" {
		fmt.Fprintf(w, "sessid param is empty")
		return
	}

	cacheData := ReadCache(sessid)
	var dct map[int]FontDot
	json.Unmarshal([]byte(cacheData), &dct)

	var data [validnum]string
	for i, v := range dct {
		data[i] = v.Text
	}

	bt, _ := json.Marshal(data)
	fmt.Fprintf(w, string(bt))
}

//验证用户点选位置正确性
func Check(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dots := r.Form.Get("dots")
	sessid := r.Form.Get("sessid")
	if dots == "" || sessid == "" {
		fmt.Fprintf(w, "dots or sessid param is empty")
		return
	}
	cacheData := ReadCache(sessid)

	src := strings.Split(dots, ",")

	var dct map[int]FontDot
	json.Unmarshal([]byte(cacheData), &dct)

	chkret := true
	for i, dot := range dct {
		j := i * 2
		k := i*2 + 1
		a, _ := strconv.Atoi(src[j])
		b, _ := strconv.Atoi(src[k])
		dist := CaluDist(a, b, dot.Dx, dot.Dy)
		if dist > iconwidth {
			chkret = false
		}

	}
	code := 1
	if chkret == true && len(dct) == validnum {
		code = 0
	}

	bt, _ := json.Marshal(map[string]interface{}{"code": code})
	fmt.Fprintf(w, string(bt))
	return
}

//计算两点距离
func CaluDist(sx, sy, dx, dy int) float64 {
	x := Abs(float64(sx - dx))
	y := Abs(float64(sy - dy))
	h := math.Hypot(x, y)
	return h
}

//直接输出验证码图片
func Outimg(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	sessid := r.Form.Get("sessid")
	if sessid == "" {
		fmt.Fprintf(w, "sessid param error")
		return
	} else {
		//图片流方式输出
		w.Header().Set("Content-Type", "image/png")

		img := GenRandImg(GenRandZh(7), sessid)
		if img == nil {
			fmt.Fprintf(w, "gen img error")
		}

		png.Encode(w, img)
	}
}

//生成随机中文图片
func GenRandImg(randstr string, sessid string) image.Image {

	if randstr == "" || sessid == "" {
		return nil
	}

	//生成画布
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// 画背景
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			//img.Set(x, y, color.RGBA{255, 255, 255, 255})
			img.Set(x, y, color.Alpha{uint8(0)})
		}
	}

	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil
	}

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	str := []rune(randstr)
	allDots := make(map[int]FontDot) //各个文字点位置
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(str); i++ {
		_w := (width - fontSize) / len(str)
		x := i*_w + rand.Intn(_w-fontSize)
		y := rand.Intn(height-fontSize) + fontSize

		pt := freetype.Pt(x, y) // 字出现的位置
		text := fmt.Sprintf("%c", str[i])
		_, err = c.DrawString(text, pt)
		if err != nil {
			return img
		}
		allDots[i] = FontDot{x, y, text}
	}

	rs := rand.Perm(len(str))
	chkDots := make(map[int]FontDot) //保存目标点位置 文字左下角
	for i, value := range rs {
		if i >= validnum {
			continue
		}
		chkDots[i] = allDots[value]
	}
	WriteCache(chkDots, sessid)

	bgpicFile := bgpicDir + (strconv.Itoa(rand.Intn(39) + 1)) + ".jpg"
	imgb, _ := os.Open(bgpicFile)
	_img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	b := _img.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m, b, _img, image.ZP, draw.Src)
	draw.Draw(m, img.Bounds(), img, image.ZP, draw.Over)
	return m
}

//写缓存
func WriteCache(v interface{}, file string) {
	if cacheType == "file" {
		bt, _ := json.Marshal(v)
		month := time.Now().Month().String()
		os.MkdirAll(cacheDir+month, 0660)
		file = cacheDir + month + "/" + file + ".json"
		logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer logFile.Close()
		io.WriteString(logFile, string(bt))
	}
}

//读缓存
func ReadCache(file string) string {
	if cacheType == "file" {
		month := time.Now().Month().String()
		file = cacheDir + month + "/" + file + ".json"
		bt, err := ioutil.ReadFile(file)
		if err == nil {
			return string(bt)
		} else {
			return ""
		}
	}
	return ""
}

//取绝对值
func Abs(x float64) float64 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}
