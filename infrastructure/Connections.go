package infrastructure

import "github.com/DiTo04/spexflix/authentication/api"

type Connections interface {
	GetAuthenticationClient() api.AuthenticationClient
}

type ConnectionsImpl struct {
	authenticationConnection api.AuthenticationClient
}

func CreateConnection(authenticationClient api.AuthenticationClient) Connections {
	return &ConnectionsImpl{authenticationConnection:authenticationClient};
}

func (c *ConnectionsImpl) GetAuthenticationClient() api.AuthenticationClient {
	return c.authenticationConnection
}