
// by Baev, 2022

package main

import (
	"net/http"
)

func hdl_CORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
}

func hdl_7001(w http.ResponseWriter, r *http.Request) {
	display(w, "403", nil)
}
func hdl_7001_upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "403", nil)
		//disp'lay(w, "upload", nil)
	case "POST":
		hdl_CORS(w)
		uplFile(w, r)
	default:
		display(w, "403", nil)
	}
}
func hdl_7001_download(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		hdl_CORS(w)
		dwlFile(w, r)
	case "POST":
		display(w, "403", nil)
	default:
		display(w, "403", nil)
	}
}

func hdl_80(w http.ResponseWriter, r *http.Request) {
	display(w, "main", nil)
}
