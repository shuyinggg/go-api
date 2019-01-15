//handlers.go contains get, post, put, and delete method handlers.
//
//Package main is restful API program
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

//Home a simple home page handler
//
func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello there!")
}

//Methods includes GET, POST, PUT and DELETE
//
func Methods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		PostCookies(w, r)
	case "POST":
		PostCookies(w, r)
	case "GET":
		GetCookies(w, r)
	case "DELETE":
		DeleteCookies(w, r)
	}
}

//GetCookies GET handlers
func GetCookies(w http.ResponseWriter, r *http.Request) {
	w = SetHeaders(w)
	json.NewEncoder(w).Encode(r.Cookies())
	logrus.Println("Get Success!")
}

//PostCookies POST and PUT handler
func PostCookies(w http.ResponseWriter, r *http.Request) {
	var c []http.Cookie
	w = SetHeaders(w)
	body, _ := ioutil.ReadAll(r.Body)
	if CheckValidInput(body) {
		json.Unmarshal(body, &c)
		if CheckDuplicate(c) {
			errmsg := `[{"err": "true", "msg": "Duplicate! Please check input"}]`
			io.WriteString(w, errmsg)
			logrus.Println("Duplicate! Please check input")
		} else {
			for _, item := range c {
				http.SetCookie(w, &http.Cookie{
					Name:  item.Name,
					Value: item.Value})
			}
			json.NewEncoder(w).Encode(&c)
			logrus.Println("Post Success!")
		}
	} else {
		errmsg := `[{"err": "true", "msg": "Invalid input! Please check input"}]`
		io.WriteString(w, errmsg)
		logrus.Println("Invalid input! Please check input")
	}

}

//DeleteCookies DELETE handler
func DeleteCookies(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	if cookies != nil {
		for _, item := range cookies {
			item := *item
			http.SetCookie(w, &http.Cookie{
				Name:    item.Name,
				Value:   item.Value,
				Expires: time.Unix(0, 0)})
		}
	}
	w = SetHeaders(w)
	json.NewEncoder(w).Encode(r.Cookies())
	logrus.Println("Delete Success!")
}
