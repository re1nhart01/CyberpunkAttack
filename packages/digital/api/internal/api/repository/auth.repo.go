package repository

import (
	"errors"
	"fmt"
	"github.com/cyberpunkattack/api/base"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/crypto"
	"github.com/cyberpunkattack/database"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/postgres"
	"github.com/cyberpunkattack/environment"
	"strings"
	"time"
)

type AuthRepository struct {
	*base.Repository
}

type InitialUser struct {
	username string
	email string
	phone string
	fullName string
	password string
}


func (repo *AuthRepository) CreateInitialUser(data InitialUser) error {
	serverKey := environment.GEnv().GetVariable("SERVER_KEY")
	userHash := crypto.GetSha1(serverKey, strings.Join([]string{data.email, data.username, time.Now().String()}, "@"))
	hashedPassword := crypto.HashPassword(data.password)

	newUserData := &models.UserModel{
		UserHash:    userHash,
		Username:    data.username,
		Email:       data.email,
		Phone:       data.phone,
		Password:    hashedPassword,
		Role: 	     database.ROLES_USER,
		FullName:    data.fullName,
		Temporary:   true,
		Active:      false,
	}

	if payload := postgres.DB().Get().Table(database.USERS_TABLE).Create(&newUserData); payload.Error != nil {
		return errors.New(fmt.Sprintf(inlineErrors.REGISTER_ERROR_1, payload.Error.Error()))
	}

	//Send Sendpulse email to validate

	return nil
}



func NewAuthRepository() *UserRepository {
	return &UserRepository{
	Repository:	&base.Repository{
			TableName: "users",
		},
	}
}