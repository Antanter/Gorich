package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5173"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println(err)
	}
}
