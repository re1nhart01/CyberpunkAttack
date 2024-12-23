package handlers

import (
	"fmt"
	"net/http"

	"github.com/cyberpunkattack/api/base"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/helpers"
	"github.com/gin-gonic/gin"
)

const USER_ROUTER = "user"

type IUserRepo interface {
	GetUserByUserHash(userHash string, is_me bool) (*models.CompoundUserProfile, error)
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
		"status":  http.StatusAccepted,
		"message": "Ok!",
	})
}

func (user *UserHandler) GetUserHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, map[string]any{
		"status":  http.StatusAccepted,
		"message": "Ok!",
	})
}

func (user *UserHandler) GetMyUserProfileHandler(context *gin.Context) {
	creds, ok := user.UnwrapUserData(context)
	if !ok {
		context.JSON(helpers.GiveUnauthorized())
		return
	}

	if userData, err := user.GetUserByUserHash(creds["userHash"].(string), true); err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_8, err.Error(), nil))
	} else {
		context.JSON(helpers.GiveOkResponseWithData(userData))
	}
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
