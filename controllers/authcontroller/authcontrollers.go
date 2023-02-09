package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/herizal95/go-jwt-mux/helper"
	"github.com/herizal95/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	user := json.NewDecoder(r.Body)
	if err := user.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var users models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek password valid
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "username atau password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// prosess pembuatan token jwt
	// expTime := time.Now().Add(time.Minute * 1)
	// claims := &config.JWTClaim{
	// 	Username: users.Username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		Issuer:    "alphacsoft.com",
	// 		ExpiresAt: jwt.NewNumericDate(expTime),
	// 	},
	// }

	// mendeklarasikan algoritma yang akan digunakan untuk signing

}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	user := json.NewDecoder(r.Body)
	if err := user.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Uuid = uuid.New()
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
