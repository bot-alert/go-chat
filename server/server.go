package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Ip   string
	Port int
}

func (s *Server) getAddress() string {
	return fmt.Sprint(s.Ip, ":", s.Port)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

type NewHandleFunc interface {
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	err := http.ListenAndServe(s.getAddress(), mux)
	if err != nil {
		panic("Count not start the server :" + err.Error())
	}
}
