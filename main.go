package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("/Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "GentlemanLin"
	indexData.Desc = "start"
	jsonstr, _ := json.Marshal(indexData)
	w.Write(jsonstr)
}

func main() {
	//	web http
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
