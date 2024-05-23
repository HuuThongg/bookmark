package api

import (
	"fmt"
	"go-bookmark/util"
	"log"
	"net/http"
)

func (h *BaseHandler) Example(w http.ResponseWriter, r *http.Request) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("cannot load config: %v", err)
		return
	}
	log.Println("jdsadsad")
	port := config.PORT
	log.Println(port)
	fmt.Fprintf(w, "heldsalo")
	fmt.Fprintf(w, "port: %v", port)
	fmt.Fprintf(w, "config\n : %v", config)

}
