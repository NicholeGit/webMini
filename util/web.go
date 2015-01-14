package util

import "net/http"

var mainServer = NewServer()

func Run(addr string) {
	mainServer.Run(addr)
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mainServer.HandleFunc(pattern, handler)
}
