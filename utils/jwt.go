package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/herizal95/hisabia_api/models/entity"
	"github.com/joho/godotenv"
)

type MyCustomeClaims struct {
	Uid      uint
	Name     string
	Username string
	Email    string
	Password string
	jwt.RegisteredClaims
}

func CreateToken(user *entity.User) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on Loading .env file")
	}

	secretKey := []byte(os.Getenv("SECKRET_KEY"))

	claims := MyCustomeClaims{
		user.Uid,
		user.Name,
		user.Username,
		user.Email,
		user.Passowrd,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on Loading .env file")
	}

	secretKey := []byte(os.Getenv("SECKRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomeClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomeClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	} else {
		fmt.Println(err)
	}
	return claims, nil
}

// func RevokeToken(tokenString string) error {
// 	// Decode token JWT
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("secret"), nil // Ganti "secret" dengan secret key yang sama dengan saat membuat token
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Ambil nilai claim "sub" dari token JWT
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return errors.New("Invalid token claims")
// 	}
// 	userID, ok := claims["sub"].(float64)
// 	if !ok {
// 		return errors.New("Invalid user ID")
// 	}

// 	// Update status token menjadi tidak valid di database
// 	_, err = config.DB.Table("users").Exec("UPDATE users SET is_valid = false WHERE id = $1", int(userID))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
