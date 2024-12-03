package handlers

import (
	"fmt"
	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/api/dtos"
	"github.com/gin-gonic/gin"
)

const AUTH_ROUTER = "auth"

type IAuthRepo interface {

}


type AuthHandler struct {
	*base.Handler
	IAuthRepo
}


func (auth *AuthHandler) GetName() string {
	return auth.Name
}

func (auth *AuthHandler) GetPath() string {
	return auth.Path
}


func (auth *AuthHandler) SignUpHandler(context *gin.Context) {
	body, stopped := auth.Unwrap(context, dtos.InitialSignUpDto)
	if stopped { return }

	fmt.Println(body)
}

func (auth *AuthHandler) ValidatePhoneOrEmailHandler(context *gin.Context) {

}


func (auth *AuthHandler) LogInHandler(context *gin.Context) {

}

func (auth *AuthHandler) RefreshTokenHandler(context *gin.Context) {
	
}




func NewAuthHandler(basePath string, repo IAuthRepo) *AuthHandler {
	return &AuthHandler{
		&base.Handler{
			Name: AUTH_ROUTER,
			Path: fmt.Sprintf("/%s/auth", basePath),
		},
		repo,
	}
}