
// by Baev, 2022

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"crypto/md5"
	"encoding/hex"
	"path/filepath"
	"encoding/json"
	"strconv"
)

func uplFile(w http.ResponseWriter, r *http.Request) {
	
	if rqkeyCheck(r.FormValue("rqkey")) != true {
		uplFile_err(w, "0")
		return
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		uplFile_err(w, "0")
		return
	}
	
	file, handler, err := r.FormFile("uplFile")
	if err != nil {
		uplFile_err(w, "2")
		return
	}
	defer file.Close()
	
	hash := md5.New()
	tm := time.Now().String()
	flsz := strconv.FormatInt(handler.Size, 10)
	
	flnm_pr := tm + r.FormValue("login") + handler.Filename + flsz
	hash.Write([]byte(flnm_pr))
	flhash := hex.EncodeToString(hash.Sum(nil))
	flext := filepath.Ext(handler.Filename)
	
	dst, err := os.Create("/var/www/fldb/"+flhash+flext)
	defer dst.Close()
	if err != nil {
		uplFile_err(w, "2")
		return
	}
	
	if _, err := io.Copy(dst, file); err != nil {
		uplFile_err(w, "2")
		return
	}
	
	resp_src := map[string]string{
		"status": "1",
		"md5": flhash,
		"name": handler.Filename,
		"ext": flext,
		"size": flsz}
	resp_json, err := json.Marshal(resp_src)
	fmt.Fprintf(w, string(resp_json))
	
}

func uplFile_err(w http.ResponseWriter, status string) {
	resp_src := map[string]string{}
	if status == "0" {
		resp_src = map[string]string{
			"status": status,
			"err": "The 403 Forbidden Error!"}
	}
	if status == "2" {
		resp_src = map[string]string{
			"status": status,
			"err": "Internal Server Error!"}
	}
	resp_json, _ := json.Marshal(resp_src)
	fmt.Fprintf(w, string(resp_json))
}
