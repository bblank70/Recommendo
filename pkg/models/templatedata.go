package models

// Business is the struct containing data about businesses
//Business structs are rendered on the various genre pages
type Business struct {
	User         string  `json:"user"`
	BusinessID   string  `json:"business_id"`
	BusinessName string  `json:"name"`
	Address      string  `json:"address"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Zipcode      int     `json:"postal_code"`
	Rating       float64 `json:"stars"`
	Photopath    string  `json:"hyperlink"`
	Model        string  `json:"Model"`
}

// The Recommendatios struct will hold results returned from the Flask API
// These are recommendations from the machine learning model
type Recommendations struct {
	Business []Business `json:"result"`
}

// TemplateData is a struct that holds data instantiated from the handlers and
// passed into render such that content can be loaded dynamically on the pages.
type TemplateData struct {
	User      string
	CSRFToken string
	Message   string
	Warning   string
	Error     string
	StringMap map[string]string
	Recs      []Business
	Popular   []Business
	Drinks    []Business
	Coffee    []Business
	New       []Business
	Yelp      []Business
	DB        DashboardData
}

// TODO: WE NEED A TYPE FOR EVERY CHART/OBJECT
type Barchart struct {
	Star5 int `json:"5-Star"`
	Star4 int `json:"4-Star"`
	Star3 int `json:"3-Star"`
	Star2 int `json:"2-Star"`
	Star1 int `json:"1-Star"`
}

var Dataslice []Barchart

// type Collection struct {
// 	Ratings    []Barchart       `json:"barchart"`
// 	History    []BusinessRating `json:"recents"`
// 	Indicators []KPI            `json:"KPI"`
// }
type Recents struct {
	BusinessRating []BusinessRating `json:"recents"`
}

// type KPI struct {
// 	MAE float64
// }

// geoslice and Geoslice allow for storage of the coordinates to map with leaflet
type Geomap map[string]float64

type BusinessRating struct {
	BusinessName  string  `json:"name"`
	PlatformStars float64 `json:"stars"`
	YourRating    int64   `json:"YourRating"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Photopath     string  `json:"hyperlink"`
}

type DashboardData struct {
	Barchart Barchart
	Geo      []Geomap
	Recents  []BusinessRating
	KPI      map[string]float64
}
