package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/records", GetPayments)
	log.Fatal(http.ListenAndServe(":8003", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/data/payment_records.json")
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Failure to read, %q", html.EscapeString(r.URL.Path))
	} else {
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)

		var mapData map[string]interface{}
		json.Unmarshal(byteValue, &mapData)
		payments, err := json.MarshalIndent(mapData, "", "\t")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Failure in data encoding, %q", html.EscapeString(r.URL.Path))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(payments)
		}
	}
}