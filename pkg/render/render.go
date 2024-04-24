package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bblank70/concierge/pkg/models"
)

// Add Default Data is for Data that we need to have available at every page.
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// //currently app runs in debug mode with RenderTemplate, in production, switch to RenderCachedTemplate in handlers.go
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html",
		"./templates/empty.layout.html")

	td = AddDefaultData(td)

	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

func RenderRecTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html",
		"./templates/empty.layout.html")

	err := parsedTemplate.Execute(w, &td)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

//TODO: We're

func RenderCachedTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//create a template cache
	tc, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	// get requested template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(ok)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err = t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println("error with buffer:", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} //creates an empty map for the template pointers

	// get all of the files *.page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	// iterate through pages
	for _, page := range pages {
		name := filepath.Base(page)                    //gets the filename of each page
		ts, err := template.New(name).ParseFiles(page) // parse the template
		fmt.Println(ts)
		if err != nil {
			fmt.Println("we had an error parsing the template")
			return myCache, err
		}
		// then the layouts
		matches, err := filepath.Glob("./templates/*.layout.html")
		fmt.Println(matches)
		if err != nil {
			fmt.Println("we had an error parsing the template")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Println("we had an error parsing the template")
				return myCache, err
			}
		}

		myCache[name] = ts // adds the template to the map of templates
	}
	return myCache, nil

}
