package api

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
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

	// Check to see if file is below max size
	// Seek to the end of the file
	offset, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println("Error uploading file: ", err)
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}

	// Get the file size (which is the current offset)
	fileSize := offset
	if fileSize > maxUploadbytes {
		fmt.Println("Error uploading file, above the max size of __ bytes")
		http.Error(w, "Error uploading file, above the max size of _ bytes", http.StatusInternalServerError)
		return
	}

	// Reset the position back to the beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		// Handle error
	}

	// Check if file is compressed, then handle

	saveExtension := filepath.Ext(handler.Filename)

	switch saveExtension {
	case ".zip":

	case ".eu4":

	default:
		fmt.Println("Invalid file type: ", err)
		http.Error(w, "Invalid file Type, must me .zip or .eu4", http.StatusInternalServerError)
		return
	}

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

func InitNostradamusAPI() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/rank", rankHandler)
	return mux
}
