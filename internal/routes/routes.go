package routes

import (
	"github.com/ViniciusDSLima/AuthSystem/internal/controller"
	"github.com/ViniciusDSLima/AuthSystem/internal/middleware"
	"github.com/gorilla/mux"
)

func RegisterUser(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/user/register", userController.RegisterUserController).Methods("POST")
}

func LoginUser(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/user/login", userController.LoginController).Methods("POST")
}

func GetUsers(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/users", middleware.AuthMiddleware(userController.GetUsers).ServeHTTP).Methods("GET")
}

func RecoveryPassword(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/user/recovery-password", userController.RecoveryPassword).Methods("POST")
}

func UpdatePassword(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/user/update-password", userController.UpdatePassword).Methods("PATCH")
}
