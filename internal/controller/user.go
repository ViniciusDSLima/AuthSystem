package controller

import (
	"encoding/json"
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
	"github.com/ViniciusDSLima/AuthSystem/internal/dto"
	authusecase "github.com/ViniciusDSLima/AuthSystem/internal/usecase/auth"
	usecase "github.com/ViniciusDSLima/AuthSystem/internal/usecase/user"
	"github.com/ViniciusDSLima/AuthSystem/internal/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
	registerUser     *usecase.RegisterUserUseCase
	login            *authusecase.LoginUseCase
	listUsers        *usecase.ListUsersUseCase
	recoveryPassword *usecase.SendTokenRecoveryPasswordUseCase
	updatePassword   *usecase.UpdatePasswordUseCase
}

func NewUserController(registerUser *usecase.RegisterUserUseCase,
	login *authusecase.LoginUseCase,
	listUsers *usecase.ListUsersUseCase,
	recoveryPassword *usecase.SendTokenRecoveryPasswordUseCase,
	updatePassword *usecase.UpdatePasswordUseCase) *UserController {
	return &UserController{
		registerUser:     registerUser,
		login:            login,
		listUsers:        listUsers,
		recoveryPassword: recoveryPassword,
		updatePassword:   updatePassword,
	}
}

var validate = validator.New()

func (c *UserController) RegisterUserController(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.UserValidate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.registerUser.Register(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(true)
}

func (c *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(loginRequest); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := c.login.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) RecoveryPassword(w http.ResponseWriter, r *http.Request) {
	var recoveryPasswordRequest dto.RecoveryPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&recoveryPasswordRequest); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(recoveryPasswordRequest); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.recoveryPassword.SendTokenToEmail(recoveryPasswordRequest.Email); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.JSONResponse(w, http.StatusOK, true)

}

func (c *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var updatePasswordRequest dto.UpdatePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&updatePasswordRequest); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(updatePasswordRequest); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.updatePassword.Execute(updatePasswordRequest); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.JSONResponse(w, http.StatusOK, true)
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.listUsers.ListUsers()

	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.JSONResponse(w, http.StatusOK, users)
}
