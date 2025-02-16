package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	port, portSpecified := os.LookupEnv("PORT")
	if !portSpecified {
		log.Printf("No port specified, so default will be used")
		port = "80"
	}
	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello mtfk. You are at %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
