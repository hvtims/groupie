package mrg

import (
	"html/template"
	"net/http"
	"strconv"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if idd <= 0 || idd >= 53 {
		http.Redirect(w, r, "/404", http.StatusFound)
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
	artist, errr := fetchLocation(artist)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}
	artist, errrr := fetchDates(artist)
	if errrr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
	}

	tmpl, err := template.ParseFiles("frontend/band.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "frontend/500.html")
		return
	}

	tmpl.Execute(w, artist)
}
