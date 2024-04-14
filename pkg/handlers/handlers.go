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

func Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html") ///used for debugging templates

	// render.RenderCachedTemplates(w, "base.html") ///puts the template in production

	// render.RenderTemplate(w, "base.html", Dataslice)
	// tpl.ExecuteTemplate(w, "base.html", Dataslice)
}
