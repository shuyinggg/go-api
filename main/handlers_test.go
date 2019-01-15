//handlers_test.go tests get, put and post methods.
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"bytes"
	"strings"
	//"fmt"
)

const url = "http://localhost:4040/cookie"
 
func TestGetCookiesNull(t *testing.T) {
	v := []byte{}
	predictedRes, _ := json.Marshal(v)

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	assertEqual(t, body, predictedRes)
}

func TestPostAndGetCookies(t *testing.T) {
	sample := `[{"Name": "Cookie1", "Value": "This is cookie1."},{"Name": "Cookie2", "Value": "This is cookie2."}]`
	predictedRes, _ := json.Marshal(sample)
	//test POST:
	postreq,_ := http.NewRequest("POST", url, strings.NewReader(sample))
	postres,_ := http.DefaultClient.Do(postreq)
	defer postres.Body.Close()
	
	//test GET:
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	assertEqual(t, body, predictedRes)
}

func TestDeleteAndGetCookies(t *testing.T) {
	predictedRes, _ := json.Marshal([]byte{})
	//test DELETE:
	dreq, _ := http.NewRequest("DELETE", url, nil)
	dres,_ := http.DefaultClient.Do(dreq)
	defer dres.Body.Close()
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	assertEqual(t, body, predictedRes)
}

func TestPostDuplicateCookies(t *testing.T) {
	predictedRes, _ := json.Marshal([]byte{})
	sample := `[{"Name": "Cookie1", "Value": "This is cookie1."},{"Name": "Cookie2", "Value": "This is cookie2."}]`
	//test POST:
	postreq,_ := http.NewRequest("POST", url, strings.NewReader(sample))
	postres,_ := http.DefaultClient.Do(postreq)
	defer postres.Body.Close()

	//POST not successful, GET should be []
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	assertEqual(t, body, predictedRes)
}

func TestPostNull(t *testing.T) {
	predictedRes, _ := json.Marshal([]byte{})
	str := ""
	//test POST:
	postreq,_ := http.NewRequest("POST", url, strings.NewReader(str))
	postres,_ := http.DefaultClient.Do(postreq)
	defer postres.Body.Close()

	//POST not successful, GET should be []
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	assertEqual(t, body, predictedRes)
}

func assertEqual(t *testing.T, actualRes []byte, predictedRes []byte) {
	if bytes.Equal(actualRes ,predictedRes) {
		t.Errorf("No cookies get, got %s, expected %s", string(actualRes), string(predictedRes))
	}
}


