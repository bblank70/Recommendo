package main

import (
	"concierge/pkg/handlers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

var Resultslice []Business

func main() {

	/////TODO - switch to cached templates
	///////This was implementing the cached templates

	// var app config.AppConfig

	// tc := render.CreateTemplateCache("base.html")

	// if err != nil {
	// 	log.Fatal("cannot create Template cache")
	// }

	// app.TemplateCache = tc
	///////////
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
	var recommendations Recommendations
	json.Unmarshal(byteValue, &recommendations)

	for i := 0; i < len(recommendations.Business); i++ {
		bizrecord := Business{
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

		// fmt.Println("BusinessID: " + recommendations.Business[i].BusinessID)
		// fmt.Println("BusinessName: " + recommendations.Business[i].BusinessName)
		// fmt.Println("Address: " + recommendations.Business[i].Address)
		// fmt.Println("City: " + recommendations.Business[i].City)
		// fmt.Println("State: " + recommendations.Business[i].State)
		// fmt.Println("Zipcode: " + strconv.Itoa(recommendations.Business[i].Zipcode))
		// fmt.Println("Rating: " + strconv.FormatFloat(recommendations.Business[i].Rating, 'f', -1, 64))
		// fmt.Println("Photo: " + recommendations.Business[i].Photopath)

	}

	// these are our paths
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/ethnic", handlers.Ethnic)
	http.HandleFunc("/drinks", handlers.AlcoholAndDrinks)
	http.HandleFunc("/coffee", handlers.Coffee)
	http.HandleFunc("/popular", handlers.Popular)
	http.HandleFunc("/new", handlers.New)
	// http.HandleFunc("request", handlers.Request)

	http.HandleFunc("/recs", handlers.Recs)

	// //this starts the server
	http.ListenAndServe(":6060", nil)

}

// ////////////////////////////////////////////

// index handles the route to the home page
// func index(w http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(w, "base.html", Resultslice)
// }
