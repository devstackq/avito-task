package server

import (
	"avito/internal/config"
	"avito/internal/handler"
	"avito/internal/repository"
	"avito/internal/service"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg *config.Config
	// db     *sql.DB
	router *gin.Engine
}

func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Init() {
	gin.SetMode(a.cfg.Mode)
	a.router = gin.New()
	a.router.Use(gin.Recovery())

	a.setComponents() // register router
}

func (a *App) Run(ctx context.Context) {
	//config custom server, , run , wait gracefull shutdown
	server := &http.Server{
		Addr:           a.cfg.Port,
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("starting web server on port: ", a.cfg.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("application forced to shutdown: ", err)
	}
	log.Println("application exiting")

}
func (a *App) setComponents() {

	apiVersion := a.router.Group("/v1")

	db, err := repository.NewPostgres(a.cfg.DBConf)
	if err != nil {
		log.Println(err)
		return
	}

	err = repository.CreateTables(db)
	if err != nil {
		log.Println(err)
		return
	}

	accountRepo := repository.NewAccountRepo(db)
	currencyRepo := repository.NewCurrencyRepo(db)

	userRepo := repository.NewUserService(db)

	accountService := service.NewAccountService(accountRepo, currencyRepo)
	userService := service.NewUserService(userRepo)

	handler.SetEnpoints(apiVersion, accountService, userService)
}
