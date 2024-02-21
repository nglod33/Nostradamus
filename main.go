package main

import (
	"fmt"
	"net/http"
)

const maxUploadbytes = 10 * 1024 * 1024 // 10 MB

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// Validate request type
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get file from request
	file, handler, err := r.FormFile("saveFile")
	if err != nil {
		fmt.Println("Error uploading file: ", err)
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Check file type

	// Check if file is compressed, then handle

	// Parse file to retrieve player information (message.txt has info on how python does this)

	// Stack rank players based on information

	// Use this stack rank to generate elo based on Microsoft research paper
	// https://www.microsoft.com/en-us/research/wp-content/uploads/2007/01/NIPS2006_0688.pdf

	// Store player and ELO data in SQLite DB

	// Create endpoint to retrieve player ranking data

	// Stretch/Optional features
	// Discord Bot that can submit saves and retrieve rankings
	// Create endpoint that will scrape saves from Skanderbeg
	// Create separate ELO rankings based on checksum
	// Scale to accomodate more than just RIT Pdx server rankings

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./web/index.html")
}

func rankHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// Hello World function
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":8080", mux)
}
