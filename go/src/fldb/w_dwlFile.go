
// by Baev, 2022

package main

import (
	"net/http"
	"io/ioutil"
	"strconv"
)

func dwlFile(w http.ResponseWriter, r *http.Request) {
	
	if rqkeyCheck(r.URL.Query().Get("rqkey")) != true {
		display(w, "403", nil)
		return
	}
	
	flnm := r.URL.Query().Get("hash") + r.URL.Query().Get("ext")
	file, err := ioutil.ReadFile("/var/www/fldb/"+flnm)
	if err != nil {
		uplFile_err(w, "2")
		return
	}
	
	w.Header().Set("Accept-ranges", "bytes")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+flnm)
	w.Header().Set("Content-Length", strconv.Itoa(len(file)))
	w.WriteHeader(http.StatusOK)
	w.Write(file)
	
}
