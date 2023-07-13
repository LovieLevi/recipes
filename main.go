package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var path string
	fs := http.FileServer(http.Dir("./static/"))
	h := http.NewServeMux()
	h.Handle("/static/", http.StripPrefix("/static/", fs))
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path = "tmpls/index.html"
		fmt.Println("Serving " + path + ";")
		b, err := os.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprint(w, string(b))
	})

    fmt.Println("Booting up RECEPTIES server on localhost:9000")
	err := http.ListenAndServe(":9000", h)
	if err != nil {
		fmt.Print(err)
	}
}
