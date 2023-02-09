package config

import "github.com/golang-jwt/jwt/v4"

var JWT_key = []byte("asdfjkl12ldjfal11sj1209fkldja")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
