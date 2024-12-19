package repository

import (
	"github.com/cyberpunkattack/api/base"
)

type ServiceRepository struct {
	*base.Repository
}



func NewServiceRepository() *ServiceRepository {
	return &ServiceRepository{
	Repository:	&base.Repository{
			TableName: "users",
		},
	}
}
