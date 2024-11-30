package repository

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/environment"
	"strconv"
)

type AppRepository struct {
	*base.Repository
}


func (app *AppRepository) TakeVer(v int64) string {
	ver := environment.GEnv().GetVariable("version")

	return ver + " " + strconv.FormatInt(v, 10)
}

func NewAppRepository() *AppRepository {
	return &AppRepository{
		Repository:     &base.Repository{TableName: ""},
	}
}