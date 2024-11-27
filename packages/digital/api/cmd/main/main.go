package main

import (
	"github.com/cyberpunkattack/api"
	"github.com/cyberpunkattack/database/postgres"
	"github.com/cyberpunkattack/database/redis"
	"github.com/cyberpunkattack/environment"
)

func main() {
	environment.InitEnvironment()
	redis.New()
	postgres.New()
	PORT := environment.GEnv().GetVariable("PORT")
	if PORT == "" {
		panic("NO PORT TOOK FROM ENVIRONMENT")
	}

	app := api.NewApp(true)
	rtErr := app.Run(PORT)
	if rtErr != nil {
		panic("APPLICATION RUNNABLE ERROR")
	}
}
