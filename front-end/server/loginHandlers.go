package main

import (
	"github.com/DiTo04/spexflix/authentication/api"
	"golang.org/x/net/context"
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

func createLoginPostHandler(client api.AuthenticationClient, logger *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Print("Could not parse form, error: " + err.Error())
			http.Error(w, "Could not parse form", http.StatusNotAcceptable)
		}
		username := r.Form.Get("name")
		password := r.Form.Get("password")
		req := &api.LoginRequest{Username: username, Password: password}
		ctx := context.Background()
		rsp, err := client.Login(ctx, req)
		switch {
		case err != nil:
			logger.Print(err.Error())
			http.Error(w, "Could not authenticate", http.StatusInternalServerError)
			break
		case !rsp.IsAuthenticated:
			http.Error(w, "Invalid credentials", http.StatusNotAcceptable)
			break
		default:
			logger.Print(rsp.SessionToken)
			cookie := &http.Cookie{
				Name:     "SessionToken",
				Path:     "/",
				Secure:   false,
				HttpOnly: false,
				Value:    rsp.SessionToken,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/browse", http.StatusFound)
		}
	}
}
