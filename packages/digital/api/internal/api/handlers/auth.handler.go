package handlers

import (
	"fmt"
	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/api/dtos"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/api/repository"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/jwt"
	"github.com/gin-gonic/gin"
)

const AUTH_ROUTER = "auth"

type IAuthRepo interface {
	CreateInitialUser(transferUser repository.InitialUser) error
	SendEmailToRedis(email, name string) error
	ValidateEmail(email, code string) (bool, error)
	LogInUser(email, password string) error
	GenerateUserTokens(email string) (*repository.UserCredentials, error)
	ValidateToken(refreshToken, grantType string) (*jwt.UserClaim, error)
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

	transferUser := repository.InitialUser{
		Username: helpers.HandleStringValues(body["username"], ""),
		Email:    helpers.HandleStringValues(body["email"], ""),
		Phone:    helpers.HandleStringValues(body["phone"], ""),
		FullName: helpers.HandleStringValues(body["fullName"], ""),
		Password: helpers.HandleStringValues(body["password"], ""),
	}

	if err := auth.CreateInitialUser(transferUser); err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_0, inlineErrors.REGISTER_ERROR_2, err.Error()))
		return
	}

	if err := auth.SendEmailToRedis(transferUser.Email, transferUser.Username); err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_1, inlineErrors.REGISTER_ERROR_2, err.Error()))
		return
	}

	context.JSON(helpers.GiveOkResponse())
}

func (auth *AuthHandler) ValidatePhoneOrEmailHandler(context *gin.Context) {
	body, stopped := auth.Unwrap(context, dtos.ValidateEmailDto)
	if stopped { return }

	email := body["email"].(string)
	code := body["code"].(string)

	isCompleted, err := auth.ValidateEmail(email, code)
	if err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_2, err.Error(), nil))
		return
	}

	context.JSON(helpers.GiveOkResponseWithData(isCompleted))
}


func (auth *AuthHandler) LogInHandler(context *gin.Context) {
	body, stopped := auth.Unwrap(context, dtos.LogInDto)
	if stopped { return }

	email := helpers.HandleStringValues(body["email"], "")
	password := helpers.HandleStringValues(body["password"], "")

	if err := auth.LogInUser(email, password); err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_3, err.Error(), nil))
		return
	}

	tokens, err := auth.GenerateUserTokens(email)
	if err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_4, err.Error(), nil))
		return
	}

	context.JSON(helpers.GiveOkResponseWithData(tokens))
}

func (auth *AuthHandler) RefreshTokenHandler(context *gin.Context) {
	body, stopped := auth.Unwrap(context, dtos.RefreshTokenDto)
	if stopped { return }

	refreshToken := helpers.HandleStringValues(body["refreshToken"], "")
	grantType := helpers.HandleStringValues(body["grantType"], "")

	claims, err := auth.ValidateToken(refreshToken, grantType)
	if err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_5, err.Error(), nil))
		return
	}

	tokens, err := auth.GenerateUserTokens(claims.Email)
	if err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_6, err.Error(), nil))
		return
	}
	
	context.JSON(helpers.GiveOkResponseWithData(tokens))
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