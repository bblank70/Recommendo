package main

import (
	"net/http"

	"github.com/bblank70/concierge/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/yelp-select", handlers.Yelp)
	http.HandleFunc("/drinks", handlers.AlcoholAndDrinks)
	http.HandleFunc("/coffee", handlers.Coffee)
	http.HandleFunc("/popular", handlers.Popular)
	http.HandleFunc("/new", handlers.New)
	// http.HandleFunc("request", handlers.Request)
	http.HandleFunc("/recs", handlers.Recs)
	http.HandleFunc("/dash", handlers.Dash)

	// //this starts the server
	http.ListenAndServe(":6060", nil)
	println("Starting the server")

}
