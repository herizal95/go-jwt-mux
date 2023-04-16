package auth

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity"
	"github.com/herizal95/hisabia_api/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var userInput entity.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	var user entity.User
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseJson(w, http.StatusNotFound, "Invalid username or password", nil)
			return
		default:
			helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	}

	// validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passowrd), []byte(userInput.Passowrd)); err != nil {
		helper.ResponseJson(w, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	token, err := utils.CreateToken(&user)
	if err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Login", token)
}

func Register(w http.ResponseWriter, r *http.Request) {

	var userInput entity.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	// hashpassword
	hashPassword, err := utils.HashPassword(userInput.Passowrd)
	if err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, "Password is Not Hashed", nil)
		return
	}

	user := entity.User{
		// Uid:      uuid.New(),
		Name:     userInput.Name,
		Email:    userInput.Email,
		Username: userInput.Username,
		Passowrd: string(hashPassword),
	}

	//insert ke db
	if err := config.DB.Create(&user).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, 201, "Register Successfully", nil)

}

// func Logout(w http.ResponseWriter, r *http.Request) {
// 	// Ambil token akses dari header Authorization
// 	aksesToken := r.Header.Get("Authorization")

// 	// Cek apakah token akses kosong
// 	if aksesToken == "" {
// 		helper.ResponseJson(w, http.StatusUnauthorized, "Unauthorized", nil)
// 		return
// 	}

// 	// Hapus prefix "Bearer " dan spasi di awal
// 	aksesToken = strings.TrimPrefix(aksesToken, "Bearer ")

// 	// Revoke token akses di sisi server
// 	err := utils.RevokeToken(aksesToken)
// 	if err != nil {
// 		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	// Berikan respons berhasil
// 	helper.ResponseJson(w, http.StatusOK, "Logout berhasil", nil)
// }
