package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend received : %s %s\n", r.Method, r.URL.Path)
	})
	fmt.Printf("Backend running on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
