
// by Baev, 2022

package main

import (
	"strings"
	"io/ioutil"
)

func rqkeyCheck(rqkey string) (result bool) {
	result = false
	org_rqkey, err := ioutil.ReadFile("/var/www/rqkey")
	org_rqkey_str := strings.TrimSpace(string(org_rqkey))
	if err == nil {
		if org_rqkey_str == rqkey {
			result = true
		}
	}
	return
}
