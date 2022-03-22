package utils

import (
	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func InitTokenAuth() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}
