package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed index.html
var content embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile("index.html")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	log.Println("Server is running on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
