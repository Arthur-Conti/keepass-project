package controllers

import (
	"net/http"
	"github.com/Arthur-Conti/keepass-project/internal/adapters"
	"github.com/Arthur-Conti/keepass-project/internal/models"
	"github.com/gin-gonic/gin"
)

type userController struct {}

func NewUserController() *userController {
	return &userController{}
}

func (controller *userController) ListAllUsersHandler(ginContext *gin.Context) {
	usersService := adapters.NewUserServiceAdapter()
	users, err := usersService.ListAllUsers()
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error listing all users: " + err.Error())
		return
	}
	sendSuccess(ginContext, http.StatusOK, users)
}

func (controller *userController) GetUserByIDHandler(ginContext *gin.Context) {
	usersService := adapters.NewUserServiceAdapter()
	userID, _ := ginContext.Params.Get("id")
	user, err := usersService.GetUserByID(userID)
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error getting user by ID: " + err.Error())
		return
	}
	sendSuccess(ginContext, http.StatusOK, user)
}

func (controller *userController) CreateUserHandler(ginContext *gin.Context) {
	usersService := adapters.NewUserServiceAdapter()
	requestContent := UserInformationRequest{}
	err := ginContext.BindJSON(&requestContent)
	if err != nil {	
		sendError(ginContext, http.StatusInternalServerError, "error binding request json: " + err.Error())
		return
	}
	err = usersService.CreateUser(models.UserModel{
		FirstName: requestContent.FirstName,
		LastName: requestContent.LastName,
		Email: requestContent.Email,
		Password: requestContent.Password,
	})
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error creating user: " + err.Error())
		return
	}
	sendSuccess(ginContext, 201, "Created")
}