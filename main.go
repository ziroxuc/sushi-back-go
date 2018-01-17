package main

import (
	"net/http"
	"log"
	"./rutas"
)

func main(){
	router := rutas.NewRouter()

	server := http.ListenAndServe(":9090",router)
	log.Fatal(server)
}



