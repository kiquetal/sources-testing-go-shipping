package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sources/m1/handlers/rest"
)

func main() {

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	println(addr)
	if addr == ":" {
		addr = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", rest.HealthCheck)
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
