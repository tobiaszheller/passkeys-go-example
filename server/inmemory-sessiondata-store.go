package server

import (
	"sync"

	"github.com/go-webauthn/webauthn/webauthn"
)

type inMemSessionDataStore struct {
	sd map[string]webauthn.SessionData
	mu sync.Mutex
}

func NewInMemSessionDataStore() *inMemSessionDataStore {
	return &inMemSessionDataStore{
		sd: map[string]webauthn.SessionData{},
	}
}

func (s *inMemSessionDataStore) Get(challenge string) *webauthn.SessionData {
	s.mu.Lock()
	defer s.mu.Unlock()
	sd, ok := s.sd[challenge]
	if !ok {
		return nil
	}
	return &sd
}

func (s *inMemSessionDataStore) Add(sd webauthn.SessionData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sd[sd.Challenge] = sd
}

func (s *inMemSessionDataStore) Delete(challenge string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sd, challenge)
}
