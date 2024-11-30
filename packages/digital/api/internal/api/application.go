package api

import (
	"fmt"
	"github.com/cyberpunkattack/api/middleware"
	"github.com/cyberpunkattack/api/routes"
	"github.com/cyberpunkattack/environment"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Application struct {
	Ver      string
	ApiPath  string
	Instance *gin.Engine
}

func NewApp(withLogger bool) *Application {
	inst := &Application{
		Ver:      environment.GEnv().GetVariable("version"),
		ApiPath:  environment.GEnv().GetVariable("API_PATH"),
		Instance: gin.Default(),
	}
	inst.Instance.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	if !withLogger {
		inst.Instance = gin.New()
	}

	inst.Instance.MaxMultipartMemory = 10 << 20

	return inst
}

func (app *Application) RunDatabaseBackgroundTasks() {
	var _ sync.WaitGroup
}

func (app *Application) BindHandlers() {
	routes.RegisterHttpAppRouter(app.Instance, app.ApiPath)
	app.Instance.Use(middleware.BodyParserMiddlewareHandler)
}

func (app *Application) Run(port string) error {

	//app.BindSocket()
	app.BindHandlers()

	httpServer := &http.Server{
		Addr:           port,
		Handler:        app.Instance,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go app.RunDatabaseBackgroundTasks()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Println("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	v := <-quit
	fmt.Println(v.String())
	fmt.Println("Server closing...")

	//defer pg.GetDatabaseInstance().Eliminate()

	return nil
}
