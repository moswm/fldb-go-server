
// by Baev, 2022

package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles(
		"/var/www/go/src/fldb/tpl/main.html",
		"/var/www/go/src/fldb/tpl/403.html",
		"/var/www/go/src/fldb/tpl/upload.html"))

func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}
