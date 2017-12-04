package main

import (
	"flag"
	au "github.com/DiTo04/spexflix/authentication"
	"github.com/DiTo04/spexflix/authentication/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:31117", "The server address in the format of host:port")
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
	ctx := context.Background()
	req := &api.LoginRequest{Username: "admin", Password: "kakakaka"}
	rsp, err := auClient.Login(ctx, req)
	if err != nil {
		panic("Got error: " + err.Error())
	}
	log.Println(rsp)

}
