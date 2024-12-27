package routes

import (
	"github.com/cyberpunkattack/api/handlers"
	"github.com/cyberpunkattack/api/repository"
	"github.com/cyberpunkattack/api/routes/http"
	"github.com/cyberpunkattack/api/routes/ws"
	"github.com/cyberpunkattack/api/service"
	"github.com/gin-gonic/gin"
)

func RegisterHttpAppRouter(engine *gin.Engine, basePath string) {
	handler := handlers.NewAppHandler(basePath, repository.NewAppRepository())
	http.AppRoute(engine, handler)
}

func RegisterHttpUserRouter(engine *gin.Engine, basePath string) {
	injectable := &repository.UserInjections{Clans: *repository.NewClansRepository()}

	handler := handlers.NewUserHandler(basePath, repository.NewUserRepository(injectable))
	http.UserRoute(engine, handler)
}

func RegisterHttpAuthRouter(engine *gin.Engine, basePath string) {
	injectable := &repository.AuthInjections{User: repository.NewUserRepository(nil)}

	handler := handlers.NewAuthHandler(basePath, repository.NewAuthRepository(injectable))
	http.AuthRoute(engine, handler)
}

func RegisterHttpSessionRouter(engine *gin.Engine, basePath string) {
	// injectable := repository.InjectableStructs{User: repository.NewUserRepository()}

	handler := handlers.NewSessionHandler(basePath, repository.NewSessionRepository())
	http.SessionRoute(engine, handler)
}

func RegisterWsGatewayRouter(engine *gin.Engine, basePath string) {
	args := handlers.GatewayHandlerArgs{
		BasePath: basePath,
		Repo:     repository.NewServiceRepository(nil),
		UserRepo: repository.NewUserRepository(&repository.UserInjections{Clans: *repository.NewClansRepository()}),
		Services: service.NewGatewayService(),
	}

	handler := handlers.NewGatewayHandler(args)
	ws.GatewayWsRoute(engine, handler)
}
