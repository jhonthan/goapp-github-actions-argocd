package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Jonathan !! New"))
	})
	http.ListenAndServe(":9090", nil)
}
