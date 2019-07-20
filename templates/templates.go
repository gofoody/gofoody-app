package templates

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var templs = template.New("")

func init() {
	for _, path := range AssetNames() {
		contents, err := Asset(path)
		if err != nil {
			log.Panicf("Failed to parse template, path: %s, err: %v", path, err)
		}
		templs.New(path).Parse(string(contents))
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templs.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
