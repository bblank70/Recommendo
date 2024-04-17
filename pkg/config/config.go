package config

import "text/template"

//AppConfig is to limit potential for import cycles
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
