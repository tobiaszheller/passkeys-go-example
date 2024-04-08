package server

import (
	"sync"
)

type inMemUserStore struct {
	users map[string]*User
	mu    sync.Mutex
}

func NewInMemoryUserStore() *inMemUserStore {
	return &inMemUserStore{
		users: map[string]*User{},
	}
}

func (us *inMemUserStore) GetByName(username string) *User {
	us.mu.Lock()
	defer us.mu.Unlock()
	u, ok := us.users[username]
	if !ok {
		return nil
	}
	return u
}

func (us *inMemUserStore) GetByID(id string) *User {
	us.mu.Lock()
	defer us.mu.Unlock()
	for _, user := range us.users {
		if user.id == id {
			return user
		}
	}
	return nil
}

func (us *inMemUserStore) Upsert(u *User) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.users[u.name] = u
}
