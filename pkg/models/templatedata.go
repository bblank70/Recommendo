package models

// TemplateData holds data from handlers to the templates/pages

type Business struct {
	BusinessID   string  `json:"business_id"`
	BusinessName string  `json:"name"`
	Address      string  `json:"address"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Zipcode      int     `json:"postal_code"`
	Rating       float64 `json:"stars"`
	Photopath    string  `json:"hyperlink"`
}

type Recommendations struct {
	Business []Business `json:"result"`
}

type TemplateData struct {
	User      string
	Data      map[string]interface{}
	CSRFToken string
	Message   string
	Warning   string
	Error     string
	StringMap map[string]string
	Business  []Business
}
