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

func fetchLocation(artist Artist) (Artist, error) {
	resp, err := http.Get(artist.Locations)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artist.Location)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

var dates Dates

func fetchDates(ala string) error {
	resp, err := http.Get(ala)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		return err
	}
	return nil
}
