package main

import (
	api "Nostradamus/api"
	"net/http"
)

const maxUploadbytes = 10 * 1024 * 1024 // 10 MB

// initiate Http server and handler
func main() {
	mux := api.InitNostradamusAPI()
	http.ListenAndServe(":8080", mux)
}
