package main

import "sync"

const (
	ROLE_NONE    = iota
	ROLE_USER    = iota
	ROLE_MANAGER = iota
	ROLE_ADMIN   = iota
)

type (
	User struct {
		Username string
		password string
		Roles    []int
	}

	UserDB struct {
		users map[string]*User
		lock  sync.Mutex
	}
)

func (udb *UserDB) init() {
	udb.users = make(map[string]*User)
	udb.addUser("jane", "always!", ROLE_USER, ROLE_MANAGER)
	udb.addUser("bob", "always!", ROLE_USER)
	udb.addUser("anne", "always!", ROLE_ADMIN)
}

func (udb *UserDB) addUser(username, passwd string, roles ...int) {
	u := &User{Username: username, password: passwd, Roles: roles}

	udb.lock.Lock()
	defer udb.lock.Unlock()
	udb.users[username] = u
}

func (udb *UserDB) getUser(username string) *User {
	udb.lock.Lock()
	defer udb.lock.Unlock()
	return udb.users[username]
}
