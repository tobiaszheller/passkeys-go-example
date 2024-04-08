package server

import (
	"crypto/rand"
	"slices"

	"github.com/go-webauthn/webauthn/webauthn"
)

var _ webauthn.User = &User{}

type User struct {
	id          string
	name        string
	credentials []webauthn.Credential
}

func NewUser(name string) *User {
	return &User{
		id:   randomString(),
		name: name,
	}
}

func (u *User) WebAuthnID() []byte {
	return []byte(u.id)
}

func (u *User) WebAuthnName() string {
	return u.name
}

func (u *User) WebAuthnDisplayName() string {
	return u.name
}

func (u *User) WebAuthnIcon() string {
	return ""
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

func (u *User) AddCredential(cred webauthn.Credential) {
	// if credential exists with given id, update it.
	idx := slices.IndexFunc(u.credentials, func(credential webauthn.Credential) bool {
		return string(credential.ID) == string(cred.ID)
	})
	if idx >= 0 {
		u.credentials[idx] = cred
		return
	}
	u.credentials = append(u.credentials, cred)
}

func randomString() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
