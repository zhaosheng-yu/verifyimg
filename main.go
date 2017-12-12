package main

import (
	"./verifypic"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/verifyImg", verifypic.Outimg)
	http.HandleFunc("/verifyChk", verifypic.Check)
	http.HandleFunc("/verifyData", verifypic.GetData)
	http.HandleFunc("/hoverclick", verifypic.Hoverclick)

	http.HandleFunc("/demo", demo)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func demo(w http.ResponseWriter, r *http.Request) {
	sessid := time.Now().UnixNano() / 1e6
	t, _ := template.ParseFiles("./demo.gtpl")
	t.Execute(w, map[string]interface{}{"sessid": sessid})
}
