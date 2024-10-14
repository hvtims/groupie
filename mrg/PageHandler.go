package mrg

import (
	"html/template"
	"net/http"
	"strconv"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idd, _ := strconv.Atoi(id)
	if idd <= 0 || idd >= 53 {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "frontend/404.html")
		return
	}
	apiURL := "https://groupietrackers.herokuapp.com/api/artists/" + id

	artist, err := fetchArtist(apiURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
	}
	artist, err = fetchRelations(artist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}

	tmpl, err := template.ParseFiles("frontend/band.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}

	tmpl.Execute(w, artist)
}
