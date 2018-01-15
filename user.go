package main

import "crypto/sha1"

type User struct {
	email    string
	userHash string
}

func (u *User) ToUserHash() {
	hash := sha1.New()
	hash.Write([]byte(u.email))
	u.userHash = string(hash.Sum(nil))
}

func NewUser(email string) *User {
	user := &User{}
	user.email = email
	user.ToUserHash()
	return user
}
