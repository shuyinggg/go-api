//server.go is the main server file
package main

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

//main main server 
func main() {
	mux := http.NewServeMux()

	logrus.Println("Server Start!")

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/cookie", Methods)
	log.Fatal(http.ListenAndServe(":4040", mux))
}
