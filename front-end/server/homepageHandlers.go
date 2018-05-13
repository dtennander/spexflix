package server

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type ContentGetter interface {
	Get(token string) (content interface{}, err error)
}

func getHomePage(htmlPath string, contentGetter ContentGetter) (func(w http.ResponseWriter, r *http.Request), error) {
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
		c, err := contentGetter.Get(cookie.Value)
		if err != nil {
			http.Error(w, "Could not create webpage.", http.StatusInternalServerError)
		}
		temp.Execute(w, c)
	}, nil

}
