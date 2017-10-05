package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "3000"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}

	env := "local"
	if v := os.Getenv("UP_STAGE"); v != "" {
		env = v
	}

	log.Printf("Listening on port %s (%s env)\n", port, env)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
