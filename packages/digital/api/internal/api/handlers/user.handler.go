package handlers

import (
	"fmt"
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

const USER_ROUTER = "user"


type IUserRepo interface {
}

type UserHandler struct {
	*base.Handler
	IUserRepo
}


func (user *UserHandler) GetName() string {
	return user.Name
}


func (user *UserHandler) GetPath() string {
	return user.Path
}

func (user *UserHandler) GetUsersListHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, map[string]any{
		"status": http.StatusAccepted,
		"message": "Ok!",
	})
}

func (user *UserHandler) GetUserHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, map[string]any{
		"status": http.StatusAccepted,
		"message": "Ok!",
	})
}

func (user *UserHandler) GetMyUserProfileHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, map[string]any{
		"status": http.StatusAccepted,
		"message": "Ok!",
	})
}


func NewUserHandler(basePath string, repo IUserRepo) *UserHandler {
	return &UserHandler{
		&base.Handler{
			Name: USER_ROUTER,
			Path: fmt.Sprintf("/%s/users", basePath),
		},
		repo,
	}
}
