//handlers_test.go tests get, put and post methods.
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	//"fmt"
)

const url = "http://localhost:4040/cookie"

func TestGetCookiesNull(t *testing.T) {
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	Methods(res, req)
	body, _ := ioutil.ReadAll(res.Body)
	var actual []byte
	json.Unmarshal(body, &actual)
	assertEqualNil(t, actual)
}

func TestPostCookies(t *testing.T) {
	cookies := []map[string]string{
		{
			"Name":  "Cookie1",
			"Value": "This is cookie1",
		},
		{
			"Name":  "Cookie2",
			"Value": "This is cookie2",
		},
	}
	// convert that into json
	//test POST:
	req, _ := http.NewRequest("POST", url, strings.NewReader(cookies))
	res := httptest.NewRecorder()
	Methods(res, req)
	body, _ := ioutil.ReadAll(res.Body)
	t.Log(body)
	var actual []byte
	json.Unmarshal(body, &actual)
	assertEqual(t, body, cookies, 1)
}

func TestDeleteCookies(t *testing.T) {
	//test DELETE:
	req, _ := http.NewRequest("DELETE", url, nil)
	res := httptest.NewRecorder()
	Methods(res, req)
	body, _ := ioutil.ReadAll(res.Body)
	var actual []byte
	json.Unmarshal(body, &actual)
	assertEqualNil(t, actual)
}

func TestPostDuplicateCookies(t *testing.T) {
	cookies := `[{"Name":"Cookie1","Value":"This is cookie1"},{"Name":"Cookie1","Value":"This is cookie2"}]`
	var sample []map[string]string
	sample1 := map[string]string{"err": "true", "msg": "Duplicate! Please check input"}
	sample = append(sample, sample1)
	//test POST:
	req, _ := http.NewRequest("POST", url, strings.NewReader(cookies))
	res := httptest.NewRecorder()
	Methods(res, req)
	body, _ := ioutil.ReadAll(res.Body)
	var actual []byte
	json.Unmarshal(body, &actual)
	assertEqual(t, body, sample, 2)
}

func TestPostNull(t *testing.T) {
	v := ""
	var sample []map[string]string
	sample1 := map[string]string{"err": "true", "msg": "Invalid input! Please check input"}
	sample = append(sample, sample1)
	//test POST:
	req, _ := http.NewRequest("POST", url, strings.NewReader(v))
	res := httptest.NewRecorder()
	Methods(res, req)
	body, _ := ioutil.ReadAll(res.Body)
	var actual []byte
	json.Unmarshal(body, &actual)
	assertEqual(t, body, sample, 2)
}

func assertEqualNil(t *testing.T, actualRes []byte) {
	if !bytes.Equal(actualRes, nil) {
		t.Errorf("Fail, got %s, expected %s", string(actualRes), `[]`)
	} else {
		t.Logf("PASS, got %s, expected %s", string(actualRes), `[]`)
	}
}

func assertEqual(t *testing.T, actualRes []byte, predictedRes []map[string]string, ty int) {
	actualSs := string(actualRes[:]) //slice then toString
	actualS := strings.Split(actualSs, "},{")
	//fmt.Println(actualS)
	var pairs []map[string]string
	for ind := range actualS {
		result := strings.FieldsFunc(actualS[ind], split)
		if ty == 1 { //cookies
			pairs = append(pairs, map[string]string{"Name": result[1], "Value": result[3]})
		} else { //errors
			length := len(result[3])
			pairs = append(pairs, map[string]string{"err": result[1], "msg": result[3][0 : length-2]})
		}

	}
	equal := true
	for ind := range pairs {
		if reflect.DeepEqual(pairs[ind], predictedRes[ind]) {
			equal = false
		}
	}
	if equal {
		t.Logf("PASS, got %v, expected %v", pairs, predictedRes)
	} else {
		t.Errorf("Fail, got %v, expected %v", pairs, predictedRes)
	}

}

func split(r rune) bool {
	return r == ':' || r == ','
}
