package routes

import "net/http"

func FindItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("received request for item " + id))
}

func FindItemLatest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("the latest is the greatest!!!"))
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created item"))
}
