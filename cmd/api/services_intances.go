package api

import (
	"github.com/ViniciusDSLima/AuthSystem/config"
	"github.com/ViniciusDSLima/AuthSystem/internal/controller"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
	"github.com/ViniciusDSLima/AuthSystem/internal/routes"
	authusecase "github.com/ViniciusDSLima/AuthSystem/internal/usecase/auth"
	usecase "github.com/ViniciusDSLima/AuthSystem/internal/usecase/user"
	"github.com/ViniciusDSLima/AuthSystem/internal/utils"
	"github.com/gorilla/mux"
)

type DependencyContainer struct {
	UserRepo                *repository.UserRepository
	TokenRepo               *repository.RecoveryTokenRepository
	TokenValidator          *utils.ValidateRecoveryTokenUseCase
	RegisterUserUsecase     *usecase.RegisterUserUseCase
	LoginUsecase            *authusecase.LoginUseCase
	ListUsersUsecase        *usecase.ListUsersUseCase
	RecoveryPasswordUsecase *usecase.SendTokenRecoveryPasswordUseCase
	UpdatePasswordUsecase   *usecase.UpdatePasswordUseCase
}

func NewDependencyContainer() (*DependencyContainer, error) {
	userRepo := repository.NewUserRepository(config.GetCollection("users"))
	tokenRepo := repository.NewRecoveryTokenRepository(config.GetCollection("token"))
	tokenValidator := utils.NewValidateRecoveryTokenUseCase(*tokenRepo)
	registerUserUsecase := usecase.NewRegisterUserUseCase(*userRepo)
	loginUsecase := authusecase.NewLoginUseCase(*userRepo)
	listUsersUsecase := usecase.NewListUsersUseCase(*userRepo)
	recoveryPasswordUsecase := usecase.NewSendTokenRecoveryPasswordUseCase(*userRepo, *tokenRepo)
	updatePasswordUsecase := usecase.NewUpdatePasswordUseCase(*userRepo, *tokenValidator)

	return &DependencyContainer{
		UserRepo:                userRepo,
		TokenRepo:               tokenRepo,
		TokenValidator:          tokenValidator,
		RegisterUserUsecase:     registerUserUsecase,
		LoginUsecase:            loginUsecase,
		ListUsersUsecase:        listUsersUsecase,
		RecoveryPasswordUsecase: recoveryPasswordUsecase,
		UpdatePasswordUsecase:   updatePasswordUsecase,
	}, nil
}

func (s *ServerApi) initializeServicesUserService(container *DependencyContainer) *mux.Router {
	userController := controller.NewUserController(
		container.RegisterUserUsecase,
		container.LoginUsecase,
		container.ListUsersUsecase,
		container.RecoveryPasswordUsecase,
		container.UpdatePasswordUsecase,
	)

	router := mux.NewRouter()

	routes.GetUsers(router, userController)
	routes.RegisterUser(router, userController)
	routes.LoginUser(router, userController)
	routes.RecoveryPassword(router, userController)
	routes.UpdatePassword(router, userController)

	return router
}
