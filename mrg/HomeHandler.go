package mrg

import (
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	apiArtist := "https://groupietrackers.herokuapp.com/api/artists"

	artists, err := fetchArtists(apiArtist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}

	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}

	tmpl.Execute(w, artists)
}
