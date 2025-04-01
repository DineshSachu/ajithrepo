package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Load IST timezone
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
		return
	}

	// Get current time in IST
	currentTime := time.Now().In(loc)
	fmt.Fprintf(w, "Current Date & Time (IST): %s", currentTime.Format("2006-01-02 15:04:05 MST"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

