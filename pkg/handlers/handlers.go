package handlers

import (
	"concierge/pkg/models"
	"concierge/pkg/render"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Global Tempalte data
var TD = models.TemplateData{}

// Recs is the handler for the /recs page;
// we'll need to get data from the model API inside this function and
// overwrite the TD.Business field

func Recs(w http.ResponseWriter, r *http.Request) {
	var Resultslice []models.Business

	// TODO: get this function to communicate with the API endpoint and replace this randomly generated content - move this to "yelp select"

	jsonFile, err := os.Open("./templates/static/response.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var recommendations models.Recommendations
	var user string

	json.Unmarshal(byteValue, &recommendations)

	for i := 0; i < len(recommendations.Business); i++ {
		bizrecord := models.Business{
			BusinessID:   recommendations.Business[i].BusinessID,
			BusinessName: recommendations.Business[i].BusinessName,
			Address:      recommendations.Business[i].Address,
			City:         recommendations.Business[i].City,
			State:        recommendations.Business[i].State,
			Zipcode:      recommendations.Business[i].Zipcode,
			Rating:       recommendations.Business[i].Rating,
			Photopath:    recommendations.Business[i].Photopath,
		}
		user = recommendations.Business[i].User

		Resultslice = append(Resultslice, bizrecord)
	}
	fmt.Println(Resultslice)

	TD.User = user
	TD.Recs = Resultslice

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)

	render.RenderRecTemplate(w, "recs.page.html", &TD)

}

// Handles request for the homepage; resets the user to empty string
func Home(w http.ResponseWriter, r *http.Request) {
	TD.User = ""
	render.RenderTemplate(w, "home.page.html", &TD) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

// Handles request for the yelp-select page
func Yelp(w http.ResponseWriter, r *http.Request) {

	//TODO: Filter down a list of the yelp select restaurants

	render.RenderTemplate(w, "yelp-select.page.html", &TD) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Coffee(w http.ResponseWriter, r *http.Request) {

	//TODO: Filter down a list of the yelp select restaurants to render in the coffee page

	render.RenderTemplate(w, "coffee.page.html", &TD) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func AlcoholAndDrinks(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "drinks.page.html", &TD) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Popular(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "popular.page.html", &TD) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func New(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new.page.html", &TD) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}
