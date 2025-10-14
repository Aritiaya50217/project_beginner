package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"id":   1,
			"name": "John Doe",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	http.ListenAndServe(":8081", nil)
}
