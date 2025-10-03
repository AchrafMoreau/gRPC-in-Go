package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodBye struct{
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye{
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Good Bye\n"));

	data, err := io.ReadAll(req.Body) 
	if err != nil{
		http.Error(res, "Oppes", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Good Bye %s\n", string(data))
}
