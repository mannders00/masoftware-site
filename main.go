package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	port := flag.String("port", "8090", "Port to run HTTP server on.")
	flag.Parse()

	tmpl := template.Must(template.New("").ParseGlob("web/templates/*"))

	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("web/public"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})
	mux.HandleFunc("/studies", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "studies.html", nil)
	})

	fmt.Println("Listening on :" + *port)
	if err := http.ListenAndServe(":"+*port, mux); err != nil {
		panic(err)
	}
}
