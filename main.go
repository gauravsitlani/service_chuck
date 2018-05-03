package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func health_check(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("OK"))
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/health_check",health_check)
}