package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/cyberpunkattack/api/middleware"
	"github.com/cyberpunkattack/api/routes"
	"github.com/cyberpunkattack/api/wstore"
	"github.com/cyberpunkattack/database"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/environment"
	"github.com/cyberpunkattack/pkg/cron"
	"github.com/gin-gonic/gin"
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

func (app *Application) RunBackgroundRoutineTasks() {
	var wg sync.WaitGroup
	ctx := context.Background()
	wg.Add(4)
	go database.CreatePostgresTables(ctx, &wg, &models.Models{})
	go database.CreatePostgresFunctions(ctx, &wg)
	go func() {
		err := wstore.ListenGlobal(wstore.AllocatedWsStore.Group.Global)
		fmt.Println("err", err)
	}()
	go func() {
		wstore.ListenSessions(wstore.AllocatedWsStore.Group.Sessions)
	}()

}

func (app *Application) TryTest() {
	ctx := context.Background()

	c := cron.New(true)
	c.CreateJob(ctx, "hubb1a", time.Duration(time.Second*25), time.Now().Add(time.Hour*5), func() error {
		fmt.Println("CRON IS EXECUTED!")
		return nil
	})
}

func (app *Application) BindSocket() {
	app.Instance.Use(middleware.BodyParserMiddlewareHandler)
	routes.RegisterWsGatewayRouter(app.Instance, app.ApiPath)
}

func (app *Application) BindHandlers() {
	routes.RegisterHttpAppRouter(app.Instance, app.ApiPath)
	app.Instance.Use(middleware.ClientBaseSecurity)
	app.Instance.Use(middleware.BodyParserMiddlewareHandler)
	routes.RegisterHttpAuthRouter(app.Instance, app.ApiPath)
	routes.RegisterHttpUserRouter(app.Instance, app.ApiPath)
	routes.RegisterHttpSessionRouter(app.Instance, app.ApiPath)
}

func (app *Application) Run(port string) error {

	app.BindSocket()
	app.BindHandlers()

	httpServer := &http.Server{
		Addr:           port,
		Handler:        app.Instance,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go app.RunBackgroundRoutineTasks()
	go app.TryTest()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Printf("Failed to listen and serve: %+v", err)
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
