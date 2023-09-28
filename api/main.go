package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Togenerate()

	fmt.Printf("listen at the port: %d", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
