
package main

import (
	"net/http"
)

func main() {
	
	mux80:=http.NewServeMux()
	mux80.HandleFunc("/",hdl_80)
	
	mux7001:=http.NewServeMux()
	mux7001.HandleFunc("/",hdl_7001)
	mux7001.HandleFunc("/upload",hdl_7001_upload)
	mux7001.HandleFunc("/download",hdl_7001_download)
	
	go http.ListenAndServe(":80",mux80)
	go http.ListenAndServe(":7001",mux7001)
	select {}
	
}
