package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/order/list", func(w http.ResponseWriter, r *http.Request) {
		resp := []map[string]interface{}{
			{"id": 101, "item": "Laptop"},
			{"id": 102, "item": "Phone"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8091", nil)
}
