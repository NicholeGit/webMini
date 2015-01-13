package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/NicholeGit/webMini/frontend"
	_ "github.com/NicholeGit/webMini/frontend/user"
)

func beforeHandle(w http.ResponseWriter, r *http.Request) bool {
	//可以做防火墙 和 访问前的一些设置
	return true
}

func afterHandle(w http.ResponseWriter, r *http.Request) {

}

type MyServeMux int

func (MyServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if beforeHandle(w, r) != true {
		return
	}
	http.DefaultServeMux.ServeHTTP(w, r)

	afterHandle(w, r)
}

func main() {

	// 静态文件
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//https://github.com/sipin/xingyun
	http.Handle("/", http.FileServer(http.Dir(wd+"\\wwwroot\\views")))
	http.Handle("/static/", http.FileServer(http.Dir(wd+"\\wwwroot\\")))

	err = http.ListenAndServe(":80", MyServeMux(0))
	if err != nil {
		log.Fatal(err)
	}
}
