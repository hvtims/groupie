package mrg

import (
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
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

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, er := template.ParseFiles("frontend/404.html")
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}
	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, nil)
}
