package main

import (
	"flag"
	au "github.com/DiTo04/spexflix/authentication"
	"github.com/DiTo04/spexflix/authentication/api"
	"google.golang.org/grpc"
	"log"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

type server struct {
	sessions au.SessionPool
}

//This server is the gateway onto Spexflix.
//
//If you are logged in you get passed to the home-page.
//Else you get the log in screen.
func main() {
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithInsecure())
	auConnection, err := grpc.Dial(*serverAddr, opt...)
	if err != nil {
		log.Fatal("Could not dial up au service, %v", err)
	}
	auClient := api.NewAuthenticationClient(auConnection)

}
