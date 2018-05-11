package main

import (
	"flag"
	"fmt"
	"github.com/DiTo04/spexflix/authentication/api"
	au "github.com/DiTo04/spexflix/authentication/authentication"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Authenticator interface {
	Login(username string, password string) (token string, err error)
	AuthenticateSession(token string) (username *string)
}

type auServer struct {
	authenticator Authenticator
	log           *log.Logger
}

func createAuService(authenticator Authenticator, log *log.Logger) *auServer {
	return &auServer{authenticator: authenticator, log: log}
}

func (s *auServer) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginReply, error) {
	s.log.Printf("Got LoginRequest: %+v", req)
	token, err := s.authenticator.Login(req.Username, req.Password)
	if err != nil {
		return &api.LoginReply{IsAuthenticated: false, SessionToken: ""}, nil
	} else {
		return &api.LoginReply{IsAuthenticated: true, SessionToken: token}, nil
	}
}

func (s *auServer) Authenticate(ctx context.Context, req *api.AuRequest) (*api.AuReply, error) {
	s.log.Printf("Got AuRequest: %+v", req)
	token := req.SessionToken
	username := s.authenticator.AuthenticateSession(token)
	if username != nil {
		return &api.AuReply{IsAuthenticated: true, Username: *username}, nil
	} else {
		return &api.AuReply{IsAuthenticated: false, Username: ""}, nil
	}
}

func main() {
	port := flag.Int64("port", 31117, "The port used for grpc connections")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
	grpcServer := grpc.NewServer()
	authenticator := au.CreateAuthenticator(au.CreateInMemorySessionPool(), nil)
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	auService := createAuService(authenticator, logger)
	api.RegisterAuthenticationServer(grpcServer, auService)
	grpcServer.Serve(lis)
	logger.Print("Starting server!")
}
