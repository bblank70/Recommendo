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
}

// TODO: Build a struct for the dashboard page.

// // RequestStruct is DEPRECATED
// type Requeststruct struct {
// 	User     string
// 	Business []Business
// }
