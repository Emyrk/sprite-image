package main

import (
	"image/png"
	"log"
	"net/http"

	"github.com/Emyrk/sprite-image/sprite"
)

func main() {
	http.HandleFunc("/generate.png", generateHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	//body := r.URL.Query().Get("body")
	//head := r.URL.Query().Get("head")
	//legs := r.URL.Query().Get("legs")
	//
	//if body == "" || head == "" || legs == "" {
	//	http.Error(w, "Missing one or more query parameters: body, head, legs", http.StatusBadRequest)
	//	return
	//}

	chacter, err := sprite.LoadSprite()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	//w.Header().Set("Content-Disposition", "attachment; filename=out.png")

	img, err := chacter.Forward()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = png.Encode(w, img)
}
