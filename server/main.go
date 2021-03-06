package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/mux"
)

// PressKey a key passed in via "key" using the keyboard in robotgo
// https://github.com/go-vgo/robotgo#keyboard
func PressKey(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	robotgo.KeyTap(params["key"])
	msg := "Pressed " + params["key"]
	json.NewEncoder(w).Encode(msg)
	log.Println(msg)
}

// Click via "Mouse" in robotgo
// https://github.com/go-vgo/robotgo#mouse
func Click(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	robotgo.MouseClick(params["key"])
	msg := "Clicked " + params["key"]
	json.NewEncoder(w).Encode(msg)
	log.Println(msg)
}

// CaptureScreen via "Bitmap" in robotgo
// https://github.com/go-vgo/robotgo#bitmap
func CaptureScreen(w http.ResponseWriter, req *http.Request) {
	bitmap := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(bitmap)
	filename := "test.png"
	robotgo.SaveBitmap(bitmap, filename)
	msg := "Captured Screen"
	json.NewEncoder(w).Encode(msg)
	log.Println(msg)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/key/{key}", PressKey).Methods("GET")
	router.HandleFunc("/click/{key}", Click).Methods("GET")
	router.HandleFunc("/capture", CaptureScreen).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}
