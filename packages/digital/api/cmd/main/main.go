package main

import (
	"github.com/cyberpunkattack/api"
	"github.com/cyberpunkattack/database/mongo"
	"github.com/cyberpunkattack/database/postgres"
	"github.com/cyberpunkattack/database/redis"
	"github.com/cyberpunkattack/environment"
	"github.com/cyberpunkattack/pkg/cron"
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/cyberpunkattack/docs"
	"github.com/swaggo/gin-swagger"
)

func init() {
	environment.New()
	redis.New()
	postgres.New()
	mongo.New()
	cron.New(true)
}

// @title           Swagger CyberpunkAttack Docs
// @version         1.0
// @description     Server Docs.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:43234
// @BasePath  /api/v2

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	PORT := environment.GEnv().Get("PORT")
	if PORT == "" {
		panic("NO PORT TOOK FROM ENVIRONMENT")
	}

	app := api.NewApp(true)
	app.Instance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rtErr := app.Run(PORT)
	if rtErr != nil {
		panic("APPLICATION RUNNABLE ERROR")
	}
}
