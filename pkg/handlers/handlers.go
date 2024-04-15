package handlers

import (
	"concierge/pkg/render"
	"net/http"
)

// index is the home page handler
func Recs(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "recs.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Ethnic(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "ethnic.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Coffee(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "coffee.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func AlcoholAndDrinks(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "drinks.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func Popular(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "popular.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}

func New(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "new.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}
