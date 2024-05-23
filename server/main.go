package main

import (
	"go-bookmark/router"
	"go-bookmark/util"
	"log"
	"net/http"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server := &http.Server{
		Addr:    config.PORT,
		Handler: router.Router(),
	}
	log.Fatal(server.ListenAndServe())

}
