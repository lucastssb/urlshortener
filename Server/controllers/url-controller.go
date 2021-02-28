package controllers

import (
	"encoding/json"
	"fmt"
	"main/db"
	"math/rand"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type urlbody struct {
	URL string
}

type respose struct {
	URLShorten string `json:"urlShorten"`
}

// GetURLHandler ...
func GetURLHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if dest, err := db.GetURLByShort(path); err == nil {
		http.Redirect(w, r, dest.URL, http.StatusFound)
	} else {
		fmt.Fprintf(w, "404 Not Found")
	}

}

// CreateURLHandler ...
func CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	var url urlbody
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}

		var randomURL = "/" + randStringBytes(6)

		shortURL := db.URL{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			URL:       url.URL,
			URLShort:  randomURL,
		}

		db.CreateURL(shortURL)

		res := respose{shortURL.URLShort}
		resposeJSON, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resposeJSON)
	} else {
		http.Error(w, "Inavalid resquest method", http.StatusMethodNotAllowed)
	}
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
