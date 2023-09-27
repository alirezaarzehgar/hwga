package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Github Actions!"))
	})
	http.ListenAndServe(":8000", nil)
}
