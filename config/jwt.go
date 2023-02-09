package config

import "github.com/golang-jwt/jwt/v4"

var JWT_key = []byte("fljalkjfoiwquoijrlknakjdshfkjh123912qjk1j2k3jlk1")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
