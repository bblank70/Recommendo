package render

import (
	"bytes"
	"concierge/pkg/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

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

// var tc = make(map[string]*template.Template)

// func RenderCachedTemplates(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]
// 	if !inMap {
// 		//create template
// 		err = CreateTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)

// 		}
// 	} else {
// 		log.Println("using cacahed template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func CreateTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

//		tc[t] = tmpl
//		return nil
//	}

// Add Default Data is for Data that we need to have available at every page.

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
