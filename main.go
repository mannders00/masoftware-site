package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "Port to run HTTP server on.")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("resources/index.html")
		if err != nil {
			panic(err)
		}
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	fmt.Println("Listening on :" + *port)
	if err := http.ListenAndServe(":"+*port, mux); err != nil {
		panic(err)
	}
}
