package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello World !");
	
	d, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Not Found", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Welcom Back %s \n", string(d))
}
