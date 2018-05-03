package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

type Request struct {
	Port 		string `json:"port"`
	Healthcheck string `json:"health_check"`
	Services	ServiceMap `json:"services"`
}

type GilmourTopic string

type Service struct{
	Group        string          `json:"group"`
	Path         string          `json:"path"`
	Timeout      int             `json:"timeout"`
	Data         interface{}     `json:"data"`
}

type ServiceMap map[GilmourTopic]Service


func health_check(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("OK"))
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/health_check",health_check)
	log.Println("Client service started .....")
	//send request here
	//Taking input based decisions from here

	http.ListenAndServe(":8088",r)
}