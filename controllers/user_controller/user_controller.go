package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity"
	"github.com/herizal95/hisabia_api/utils"
)

func MyProfile(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("userinfo").(*utils.MyCustomeClaims)

	userResponse := &entity.User{
		Uid:      user.Uid,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Passowrd: user.Password,
	}

	helper.ResponseJson(w, 200, "My Profile", userResponse)
}

func GetUserID(w http.ResponseWriter, r *http.Request) {

	uid := r.URL.Query().Get("uid")

	var user entity.User

	if err := config.DB.Where("uid = ?", uid).First(&user).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "User not found", nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get User by ID", user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	uid := r.URL.Query().Get("uid")

	var user entity.User
	if err := config.DB.Where("uid = ?", uid).First(&user).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var userUpdate entity.User
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashPassword, err := utils.HashPassword(userUpdate.Passowrd)
	if err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user.Name = userUpdate.Name
	user.Username = userUpdate.Username
	user.Email = userUpdate.Email
	user.Passowrd = hashPassword

	if err := config.DB.Updates(&user).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update User", user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	helper.ResponseJson(w, 201, "Successfully to create user data", user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	uid := r.URL.Query().Get("uid")

	var user entity.User

	if err := config.DB.Where("uid = ?", uid).First(&user).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "User not found", nil)
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete user", user)
}
