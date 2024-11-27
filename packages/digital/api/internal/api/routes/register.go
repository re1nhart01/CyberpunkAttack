package routes

import (
	"github.com/cyberpunkattack/api/handlers"
	"github.com/cyberpunkattack/api/repository"
	"github.com/cyberpunkattack/api/routes/http"
	"github.com/gin-gonic/gin"
)

func RegisterHttpAppRouter(engine *gin.Engine, basePath string) {
	handler := handlers.NewAppHandler(basePath, repository.NewAppRepository())
	http.AppRoute(engine, handler)
}