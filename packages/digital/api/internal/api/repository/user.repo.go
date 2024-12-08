package repository

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/database"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/postgres"
)

type UserRepository struct {
	*base.Repository
}


func (user *UserRepository) GetUserByEmail(email string) (*models.UserModel, bool) {
	userModel := models.UserModel{}
	payload := postgres.DB().Get().Table(database.USERS_TABLE).Where("email = ?", email).First(&userModel)
	if payload.Error != nil {
		return &userModel, false
	}

	return &userModel, true
}


func NewUserRepository() *UserRepository {
	return &UserRepository{
	Repository:	&base.Repository{
			TableName: "users",
		},
	}
}