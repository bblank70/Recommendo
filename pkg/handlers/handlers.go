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

// index is the home page handler
func Recs(w http.ResponseWriter, r *http.Request) {
	var Resultslice []models.Business

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
		Resultslice = append(Resultslice, bizrecord)
	}
	fmt.Println(Resultslice)
	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)

	render.RenderRecTemplate(w, "recs.page.html", &Resultslice)

}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{}) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Ethnic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "ethnic.page.html", &models.TemplateData{}) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Coffee(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "coffee.page.html", &models.TemplateData{}) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func AlcoholAndDrinks(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "drinks.page.html", &models.TemplateData{}) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Popular(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "popular.page.html", &models.TemplateData{}) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func New(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new.page.html", &models.TemplateData{}) ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}
func Request(w http.ResponseWriter, r *http.Request) {

	// params := url.Values{}
	// params.Add("body", "This is a test message")

	// params.Add("id", "1")

	// resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", params)
	// if err != nil {
	// 	log.Printf("Request Failed: %s", err)
	// 	return
	// }

	// data := make(map[string]interface{})
	// data["Data"] = resp

	// fmt.Println(resp)

	render.RenderTemplate(w, "request.page.html", &models.TemplateData{
		// Data: data,
	})

}
