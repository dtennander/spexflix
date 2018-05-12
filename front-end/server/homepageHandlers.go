package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type content struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func getHomePage(htmlPath string, contentServerAdress string) (func(w http.ResponseWriter, r *http.Request), error) {
	data, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		log.Print("Could not read file: " + htmlPath)
		log.Print("Error: " + err.Error())
		return nil, err
	}
	temp, err := template.New("homepage").Parse(string(data))
	if err != nil {
		log.Print("Could not parse data in homepage.tmpl")
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("SessionToken")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		log.Print("Found cookie:" + cookie.Value)
		targetUri := "http://" + contentServerAdress + "/content/" + cookie.Value
		log.Print("Getting: " + targetUri)
		js, err := http.Get(targetUri)
		if err != nil {
			http.Error(w, "Could not find content server. Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer js.Body.Close()
		c := &content{}
		err = json.NewDecoder(js.Body).Decode(c)
		if err != nil {
			http.Error(w, "Could not decode json. Error: "+err.Error(), http.StatusInternalServerError)
		}
		temp.Execute(w, c)
	}, nil

}
