//cookies_handler.go contains functions that parsing  the input cookies
package main

import(
	"net/http"
	"encoding/json"
)


//SetHeaders sets the header Cotent-tYPE for cookies
func SetHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; encoding=utf-8")
	w.WriteHeader(http.StatusOK)
	return w
}

//CheckValidInput checks if the input is valid JSON
func CheckValidInput(body []byte) bool{
	return json.Valid(body) 
}

//CheckDuplicate checks if the input is duplicate
func CheckDuplicate(c []http.Cookie) bool {
	set := make(map[string]struct{})
	for _,item := range c {
		_, ok:= set[item.Name]
		if !ok {
			set[item.Name] = struct{}{}
		} else {
			return true
		}
	}
	return false
}




