package server

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func createLoginGetHandler(logger *log.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("/html-templates/login-page.tmpl")
		if err != nil {
			logger.Fatal("Got error: " + err.Error())
		}
		tmpl, err := template.New("Logi").Parse(string(file))
		if err != nil {
			logger.Fatal("Got error: " + err.Error())
		}
		tmpl.Execute(w, nil)
	}
}

type Loginer interface {
	Login(username string, password string) (token string, err error)
}

func createLoginPostHandler(auClient Loginer, logger *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Print("Could not parse form, error: " + err.Error())
			http.Error(w, "Could not parse form", http.StatusNotAcceptable)
		}
		username := r.Form.Get("name")
		password := r.Form.Get("password")
		token, err := auClient.Login(username, password)
		if err != nil {
			logger.Print(err.Error())
			http.Error(w, "Could not authenticate", http.StatusInternalServerError)
		}
		cookie := &http.Cookie{
			Name:     "SessionToken",
			Path:     "/",
			Secure:   false,
			HttpOnly: false,
			Value:    token,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/browse", http.StatusFound)
	}
}
