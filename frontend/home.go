package frontend

import (
	"io"
	"io/ioutil"
	"net/http"
)

func init() {
	//http.HandleFunc("/", HomeHandler)
}

// 首页
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	filename := "views/login.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		io.WriteString(w, err.Error())
	}
	w.Write(body)
	return
}
