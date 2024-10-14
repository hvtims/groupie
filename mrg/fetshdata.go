package mrg

import (
	"encoding/json"
	"net/http"
)

func fetchArtist(apiURL string) (Artist, error) {
	var artist Artist
	resp, err := http.Get(apiURL)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

func fetchArtists(apiURL string) (Artists, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists Artists
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

// func fetchLocations(apiURL string) (location, error) {
// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var location location
// 	err = json.NewDecoder(resp.Body).Decode(&location)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return location, nil
// }

// func fetchDates(apiURL string) (date, error) {
// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var date date
// 	err = json.NewDecoder(resp.Body).Decode(&date)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return date, nil
// }

func fetchRelations(artist Artist) (Artist, error) {
	resp, err := http.Get(artist.Relations)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&artist.Relation)
	if err != nil {
		return artist, err
	}
	return artist, nil
}
