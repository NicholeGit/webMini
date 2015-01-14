package util

import (
	"log"
	"net/http"
	"os"
)

func beforeHandle(w http.ResponseWriter, r *http.Request) bool {
	//可以做防火墙 和 访问前的一些设置
	return true
}

func afterHandle(w http.ResponseWriter, r *http.Request) {

}

type Server struct {
	mux *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if beforeHandle(w, r) != true {
		return
	}
	s.mux.ServeHTTP(w, r)

	afterHandle(w, r)
}

func (s *Server) Run(addr string) {
	s.initServer()

	// 静态文件
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	s.mux.Handle("/", http.FileServer(http.Dir(wd+"\\wwwroot\\views")))
	s.mux.Handle("/static/", http.FileServer(http.Dir(wd+"\\wwwroot\\")))

	err = http.ListenAndServe(addr, s.mux)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) initServer() {
	//if s.Config == nil {
	//	s.Config = &ServerConfig{}
	//}

	//if s.Logger == nil {
	//	s.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//}
}

func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}
