package main

import (
	"context"
	"github.com/DiTo04/spexflix/authentication/api"
	"github.com/DiTo04/spexflix/infrastructure"
	"google.golang.org/grpc"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	serverAddr = os.Getenv("AUTHENTICATION_SERVER")
	auPort = os.Getenv("AUTHENTICATION_PORT")
	logger = log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
)

type server struct {
	connections infrastructure.Connections
	logger      *log.Logger
}

func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/login":
		s.HandleLogin(rw, r)
	default:
		s.logger.Print("Got: " + r.RequestURI)
	}
}
func (s *server) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		file, err := ioutil.ReadFile("gateway/login-page.tmpl")
		if err != nil {
			s.logger.Fatal("Got error: " + err.Error())
		}
		tmpl, err := template.New("Logi").Parse(string(file))
		if err != nil {
			s.logger.Fatal("Got error: " + err.Error())
		}
		tmpl.Execute(rw, nil)
	case "POST":
		err := r.ParseForm()
		if err != nil {
			s.logger.Print("Could not parse form, error: " + err.Error())
			http.Error(rw, "Could not parse form", http.StatusNotAcceptable)
		}
		username := r.Form.Get("name")
		password := r.Form.Get("password")
		req := &api.LoginRequest{Username: username, Password: password}
		ctx := context.Background()
		rsp, err := s.connections.GetAuthenticationClient().Login(ctx, req)
		switch {
		case err != nil:
			s.logger.Print(err.Error())
			http.Error(rw, "Could not authenticate", http.StatusInternalServerError)
			break;
		case !rsp.IsAuthenticated:
			http.Error(rw, "Invalid credentials", http.StatusNotAcceptable)
			break;
		default:
			s.logger.Print(rsp.SessionToken)
			cookie := &http.Cookie{
				Name:     "SessionToken",
				Path:     "/",
				Secure:   false,
				HttpOnly: false,
				Value:    rsp.SessionToken,
			}
			http.SetCookie(rw, cookie)
			http.Redirect(rw, r, "/browse", http.StatusFound)
		}

	}
}

//This server is the gateway onto Spexflix.
//If you are logged in you get passed to the home-page.
//Else you get the log in screen.
func main() {
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithInsecure())
	addrAndPort := serverAddr + ":" + auPort
	logger.Print("Connecting to: " + addrAndPort)
	auConnection, err := grpc.Dial(addrAndPort, opt...)
	if err != nil {
		log.Fatal("Could not dial up au service, %v", err)
	}
	auClient := api.NewAuthenticationClient(auConnection)
	connections := infrastructure.CreateConnection(auClient)
	server := &server{connections: connections, logger: logger}
	http.ListenAndServe("0.0.0.0:8000", server)
}
