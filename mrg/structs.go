package mrg

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations	string   `json:"relations"`
	Relation	Relation   
}
type Artists []Artist

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Name      string   `json:"name"`
	Dates     string   `json:"dates"`
}

// type location []Locations

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
// type date []Dates

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type All struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
