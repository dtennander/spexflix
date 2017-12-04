package main

import (
	"flag"
	"fmt"
	au "github.com/DiTo04/spexflix/authentication"
	"github.com/DiTo04/spexflix/authentication/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type auServer struct {
	authenticator au.Authenticator
	sessions      au.SessionPool
	log           *log.Logger
}

func createAuService(authenticator au.Authenticator, sessionPool au.SessionPool, log *log.Logger) *auServer {
	return &auServer{authenticator: authenticator, sessions: sessionPool, log: log}
}

func (s *auServer) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginReply, error) {
	s.log.Printf("Got LoginRequest: %+v", req)
	username := req.Username
	password := req.Password
	success := s.authenticator.Authenticate(username, password)
	rsp := &api.LoginReply{IsAuthenticated: success}
	if success {
		session := s.sessions.CreateSession(username)
		rsp.SessionToken = session.GetSessionId()
	}
	return rsp, nil
}

func (s *auServer) Authenticate(ctx context.Context, req *api.AuRequest) (*api.AuReply, error) {
	s.log.Printf("Got AuRequest: %+v", req)
	token := req.SessionToken
	isValid := s.sessions.IsSessionIdValid(token)
	username, err := s.sessions.GetUsername(token)
	if err != nil || !isValid {
		username = ""
	}
	return &api.AuReply{IsAuthenticated: isValid, Username: username}, nil
}

func main() {
	port := flag.Int64("port", 31117, "The port used for grpc connections")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	authenticator := &au.AuthenticatorImpl{}
	sp := au.SessionPoolImpl{}
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	auService := createAuService(authenticator, sp, logger)
	api.RegisterAuthenticationServer(grpcServer, auService)
	grpcServer.Serve(lis)
}
