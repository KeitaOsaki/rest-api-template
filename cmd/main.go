package main

import (
	"api_template/app/configs"
	"api_template/app/interfaces"

	"log"
	"net/http"
)

func main() {

	r := interfaces.NewServer()
	err := r.Init()
	if err != nil {
		log.Fatal(err)
	}
	addr := configs.GetServerPort()
	log.Println("port", addr, "Starting app")
	http.ListenAndServe(addr, r.Router)
}
