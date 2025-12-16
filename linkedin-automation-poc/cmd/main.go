package main

import (
	"log"
	"net/http"
)

func main() {
    // Serve root html files
    http.Handle("/", http.FileServer(http.Dir(".")))

    log.Println("Server running at http://localhost:8080/")
    err := http.ListenAndServe(":9090", nil)

    if err != nil {
        log.Fatal(err)
    }
}
