package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main(){
	server:=&http.Server{
		Addr:":3000",
		Handler: http.HandlerFunc(myHandler),
	}
	err:=server.ListenAndServe()
	if err!=nil{
		fmt.Println("Error listening to server.")
	}
} 

func myHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hi guy"))
}