<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"/>
<title>点选图片验证码</title>
<script src="http://libs.baidu.com/jquery/2.1.1/jquery.min.js"></script>
</head>

<style>
    body{margin:0;padding:0;}
    #wp{position:relative;}
</style>

<body>
<div id="wp">
    <img id="imgdiv" src="/verifyImg?sessid={{.sessid}}" onclick="appendicon(event);"/>
    <input type="hidden" value="" id="dots"/>
</div>
<div><p>>>>请依次点击 <span id="words"></span>完成验证>>></p></div>
<div><p style="margin-left:50px;"><button onclick="checkdot();">点击验证</button></p></div>
<div><p style="margin-left:50px;" id="prom"></p></div>
</body>

<script type="text/javascript"> 
//验证码图片大小文 320x100, icon大小为22x22
function appendicon(event) {
    var x=event.clientX;
    var y=event.clientY;
    var _x = x - 11;
    var _y = y - 11;
    var icon = "<img src='/hoverclick' style='position:absolute;top:"+ _y +"px;left:"+ _x +"px;'/>"
    $("#wp").append(icon);

    var dot = $("#dots").val();
    if (dot == '') {
        $("#dots").val(_x +','+ (y+11)); //传icon图左下角点坐标
    } else {
        $("#dots").val(dot + ','+ _x +','+ (y+11));
    }
}

function checkdot() {
    var dots = $("#dots").val();
    $.ajax({
            url: '/verifyChk',
            dataType: 'json',
            type: 'GET',
            data: {'dots': dots, 'sessid':{{.sessid}}},
            success: function(data){
                if (data.code == '0') {
                    $("#prom").text('验证成功,^_^!');
                } else {
                    $("#prom").text('验证失败!');
                }
            },
    });
}
$(function(){
    $.ajax({
            url: '/verifyData',
            dataType: 'json',
            type: 'GET',
            data: {'sessid':{{.sessid}}},
            success: function(data){
            console.log(data);
                //data = eval("("+data+")");
                var words = '';
                for (i=0; i<data.length; i++) {
                    words += ' "'+ data[i] +'" ';
                }
                $("#words").text(words);
            },
    });
});
</script> 

</html>
