package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/date", func(w http.ResponseWriter, r *http.Request) {
		dates := make([]struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		}, 10000)

		currentTime := time.Now()
		for i := range dates {
			dates[i] = struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			}{
				Year:  currentTime.Year(),
				Month: int(currentTime.Month()),
				Day:   currentTime.Day(),
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct{ Data interface{} }{Data: dates})
	})

	http.HandleFunc("/api/name", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
			return
		}

		var names []string
		for i := 0; i < 10000; i++ {
			names = append(names, name)
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(struct{ Data interface{} }{Data: names})
		if err != nil {
			return
		}
	})

	fmt.Println("Сервер запущен на порту :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
