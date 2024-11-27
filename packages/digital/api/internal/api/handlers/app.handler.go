package handlers

import (
	"fmt"
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


const APP_ROUTER = "app"

type IAppRepo interface {
	TakeVer(now int64) string
}

type AppHandler struct {
	*base.Handler
	IAppRepo
}

func (app *AppHandler) GetName() string {
	return app.Name
}

func (app *AppHandler) GetPath() string {
	return app.Path
}

type GetHandlerResponse struct {
    IsAlive bool   `json:"isAlive"`
    Ver     string `json:"ver"`
}

// GetHandler provides the API status and version.
// @Summary Check API health and version
// @Description Returns a JSON response indicating the API is alive and the current version based on the timestamp.
// @Tags health
// @Produce json
// @Success 200 {object} GetHandlerResponse "Response containing isAlive status and version"
// @Router / [get]
func (app *AppHandler) GetHandler(context *gin.Context) {
	ver := app.IAppRepo.TakeVer(time.Now().Unix())

	context.JSON(http.StatusOK, map[string]any{
		"isAlive": true,
		"ver": ver,
	})
}


func NewAppHandler(basePath string, repo IAppRepo) *AppHandler {
	return &AppHandler{
		&base.Handler{
			Name: APP_ROUTER,
			Path: fmt.Sprintf("/%s/", basePath),
		},
		repo,
	}
}