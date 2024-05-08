package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/bblank70/concierge/pkg/models"
	"github.com/bblank70/concierge/pkg/render"
)

// Global Tempalte data
var TD = models.TemplateData{}

// endpoints for local debug
var baseURL = "http://127.0.0.1:5000"
var reqresource = "/returnjson"
var dashresource = "/dashboard"

//endpoints for cloud production
// var baseURL = "https://recommendoapi-pitlxbolsa-uc.a.run.app/"
// var resource = "/returnjson"

// Recs is the handler for the /recs page;
// we'll need to get data from the model API inside this function and
// overwrite the TD.Business field
func Recs(w http.ResponseWriter, r *http.Request) {

	var Resultslice []models.Business
	// jsonFile, err := os.Open("./templates/static/response.json")
	// TODO: get this function to communicate with the API endpoint and replace this randomly generated content - move this to "yelp select"

	user := r.FormValue("User")
	if user != "" {
		TD.User = user
	}

	params := url.Values{}
	params.Add("user", TD.User)

	u, _ := url.ParseRequestURI(baseURL)
	fmt.Println("The uRL was:", u)
	u.Path = reqresource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	resp, err := http.Get(urlStr)

	// resp, err := http.Get(posturl)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	fmt.Println(resp)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("There was an error recieving the body:", err)
	}

	var recommendations models.Recommendations

	json.Unmarshal(b, &recommendations)

	for i := 0; i < len(recommendations.Business); i++ {
		bizrecord := models.Business{
			User:         recommendations.Business[i].User,
			BusinessID:   recommendations.Business[i].BusinessID,
			BusinessName: recommendations.Business[i].BusinessName,
			Address:      recommendations.Business[i].Address,
			City:         recommendations.Business[i].City,
			State:        recommendations.Business[i].State,
			Zipcode:      recommendations.Business[i].Zipcode,
			Rating:       recommendations.Business[i].Rating,
			Photopath:    recommendations.Business[i].Photopath,
		}

		Resultslice = append(Resultslice, bizrecord)
	}
	fmt.Println(Resultslice)

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

func Dash(w http.ResponseWriter, r *http.Request) {

	// var Recentslice []models.Recents

	user := TD.User
	if user != "" {
		TD.User = user
	}

	params := url.Values{}
	params.Add("user", TD.User)

	u, _ := url.ParseRequestURI(baseURL)
	fmt.Println("The uRL was:", u)
	u.Path = dashresource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	resp, err := http.Get(urlStr)

	// resp, err := http.Get(posturl)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	fmt.Println(resp)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("There was an error recieving the body:", err)
	}

	var dbdata models.DashboardData

	json.Unmarshal(b, &dbdata)

	for i := 0; i < len(dbdata.Recents); i++ {
		bizrecord := models.BusinessRating{
			BusinessName:  dbdata.Recents[i].BusinessName,
			PlatformStars: dbdata.Recents[i].PlatformStars,
			YourRating:    dbdata.Recents[i].YourRating,
			Photopath:     dbdata.Recents[i].Photopath,
		}
		fmt.Println(bizrecord)
		// 	Recentslice = append(Recentslice, bizrecord)
	}
	// fmt.Println(Recentslice)

	//	TD.DB = Resultslice

	render.RenderRecTemplate(w, "dash.page.html", &TD)

}
