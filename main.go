package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
func main() {
	res, err := http.Get("https://swapi.co/api/people/1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	luke := &Luke{}
	err = json.Unmarshal(body, &luke)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))
	http.HandleFunc("/luke", luke.printer)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
type Luke struct {
	Name string `json:"name"`
	Height string   `json:"height"`
}
func (l *Luke) printer(w http.ResponseWriter, r *http.Request) {
	stringLuke, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	w.Write(stringLuke)
}

