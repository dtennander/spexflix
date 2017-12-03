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
)

type auServer struct {
	authenticator au.Authenticator
	sessions      au.SessionPool
}

func createAuService(authenticator au.Authenticator, sessionPool au.SessionPool) *auServer {
	return &auServer{authenticator: authenticator, sessions: sessionPool}
}

func (s *auServer) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginReply, error) {
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
	token := req.SessionToken
	isValid := s.sessions.IsSessionIdValid(token)
	username, err := s.sessions.GetUsername(token)
	if err != nil || !isValid {
		username = ""
	}
	return &api.AuReply{IsAuthenticated: isValid, Username: username}, nil
}

func main() {
	port := flag.Int64("port", 1337, "The port used for grpc connections")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	authenticator := &au.AuthenticatorImpl{}
	sp := au.SessionPoolImpl{}
	auService := createAuService(authenticator, sp)
	api.RegisterAuthenticationServer(grpcServer, auService)
	grpcServer.Serve(lis)
}
