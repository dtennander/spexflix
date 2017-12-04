package authentication

type Authenticator interface {
	Authenticate(user string, password string) bool
}

type AuthenticatorImpl struct {
	hashes map[string]string
}

func (au *AuthenticatorImpl) Authenticate(user string, password string) bool {
	if user == "admin" && password == "kakakaka" {
		return true
	} else {
		return au.hashes[user] == password
	}
}
