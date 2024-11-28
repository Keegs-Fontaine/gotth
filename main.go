package main

import (
	"fmt"
	"net/http"

	"github.com/Keegs-Fontaine/gotth/src/views"
	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	component := views.Index()
	mux.Handle("/", templ.Handler(component))

	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./tmp/styles"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
