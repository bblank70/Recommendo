package main

import (
	"concierge/pkg/handlers"
	"net/http"
)

func main() {

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
