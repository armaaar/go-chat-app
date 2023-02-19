package main

import (
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func healthRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "{\"state\":\"ok\"}")
}

func main() {
	http.HandleFunc("/", healthRoute)

	log.Panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
