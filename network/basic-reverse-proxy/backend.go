package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend received : %s %s\n", r.Method, r.URL.Path)
	})
	fmt.Println("Backend running on port : 8081")
	http.ListenAndServe(":8081", nil)
}
