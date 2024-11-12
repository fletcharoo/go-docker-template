package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fletcharoo/snest"
)

type Config struct {
	Port string `snest:"PORT"`
}

func main() {
	var conf Config
	if err := snest.Load(&conf); err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	http.HandleFunc("/", response)
	addr := ":" + conf.Port
	log.Printf("Serving on %s", addr)
	http.ListenAndServe(addr, nil)
}

func response(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
