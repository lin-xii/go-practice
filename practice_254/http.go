package practice254

import "net/http"

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/ index"))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/hello handler"))
	})
	http.ListenAndServe(":8080", nil)
}
