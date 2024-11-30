package repository

import "github.com/cyberpunkattack/api/base"

type AuthRepository struct {
	*base.Repository
}

func NewAuthRepository() *UserRepository {
	return &UserRepository{
	Repository:	&base.Repository{
			TableName: "users",
		},
	}
}