package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ys", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Host:", r.Host)
		w.Write([]byte("原神"))
	})
	http.ListenAndServe(":8080", nil)
}
