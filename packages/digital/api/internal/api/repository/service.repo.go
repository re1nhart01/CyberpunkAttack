package repository

import (
	"github.com/cyberpunkattack/api/base"
)

type ServiceRepository struct {
	*base.Repository
	injectable *ServiceInjectables
}

type ServiceInjectables struct {
}

func NewServiceRepository(injectable *ServiceInjectables) *ServiceRepository {
	return &ServiceRepository{
		Repository: &base.Repository{
			TableName: "users",
		},
		injectable: injectable,
	}
}
