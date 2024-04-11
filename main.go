package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

var tpl *template.Template

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

// type RecommendataionResult []Business

// init instantiates the templates, they must be .tmpl extenstions
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

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

		fmt.Println("BusinessID: " + recommendations.Business[i].BusinessID)
		fmt.Println("BusinessName: " + recommendations.Business[i].BusinessName)
		fmt.Println("Address: " + recommendations.Business[i].Address)
		fmt.Println("City: " + recommendations.Business[i].City)
		fmt.Println("State: " + recommendations.Business[i].State)
		fmt.Println("Zipcode: " + strconv.Itoa(recommendations.Business[i].Zipcode))
		fmt.Println("Rating: " + strconv.FormatFloat(recommendations.Business[i].Rating, 'f', -1, 64))
		fmt.Println("Photo: " + recommendations.Business[i].Photopath)

	}

	// these are our paths
	http.HandleFunc("/", index)
	// http.HandleFunc("/verify", verifyer)
	// http.HandleFunc("/response", responder)
	// //this starts the server
	http.ListenAndServe(":6060", nil)

}

// ////////////////////////////////////////////

// index handles the route to the home page
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "base.html", Resultslice)
}
