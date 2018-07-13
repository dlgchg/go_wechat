package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {

	http.HandleFunc("/wechat", getWechatServerInfo)
	error := http.ListenAndServe(":9000", nil)
	if error != nil {
		log.Fatalln("ListenAndServe: ",error)
	}
}

func getWechatServerInfo(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()

	fmt.Println(r.Form)
}
