package main

import (
	"crypto/sha1"
)

type User struct {
	email string
}

func (u *User) ToUserHash() []byte {
	hash := sha1.New()
	hash.Write([]byte(u.email))
	return hash.Sum(nil)
}

func NewUser(email string) *User {
	return &User{email}
}
