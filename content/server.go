package main

import (
	"encoding/json"
	"github.com/DiTo04/spexflix/authentication/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	serverAddr = os.Getenv("AUTHENTICATION_SERVER")
	auPort     = os.Getenv("AUTHENTICATION_PORT")
)

type content struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

const contentMessage = "Välkommen till Spexflix!"

func getContentForUser(user string) content {
	return content{Username: user, Content: contentMessage}
}

func getApiHandler(auClient api.AuthenticationClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Got Request!")
		s := strings.Split(r.RequestURI, "/")
		sessionToken := s[len(s)-1]
		ctx := context.Background()
		req := &api.AuRequest{SessionToken: sessionToken}
		rsp, err := auClient.Authenticate(ctx, req)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		if !rsp.IsAuthenticated {
			http.Error(w, "Not authenticated", http.StatusNotAcceptable)
			return
		}
		content := getContentForUser(rsp.Username)
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(true)
		encoder.Encode(content)
	}

}

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	addrAndPort := serverAddr + ":" + auPort
	logger.Print("Connecting to: " + addrAndPort)
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithInsecure())
	auConnection, err := grpc.Dial(addrAndPort, opt...)
	if err != nil {
		log.Fatal("Could not dial up au service,", err)
	}
	auClient := api.NewAuthenticationClient(auConnection)
	http.HandleFunc("/content/", getApiHandler(auClient))
	http.ListenAndServe("0.0.0.0:8000", nil)
}
