package server

import (
	"context"
	"github.com/diptomondal007/go_clean_arch/auth"
	authhttp "github.com/diptomondal007/go_clean_arch/auth/delivery/http"
	authmysql "github.com/diptomondal007/go_clean_arch/auth/repository/mysql"
	authusecase "github.com/diptomondal007/go_clean_arch/auth/usecase"
	"github.com/diptomondal007/go_clean_arch/config"
	"github.com/diptomondal007/go_clean_arch/conn"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
}

func NewApp(appCfg *config.App) *App {
	app := new(App)
	router := echo.New()

	// middleware
	router.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	db , _ := conn.GetDB()
	authRepo := authmysql.NewUserRepository(db.DB, viper.GetString("app.auth_table"))
	app.authUC = authusecase.NewAuthUseCase(authRepo, []byte(viper.GetString("auth.signing_key")), viper.GetDuration("auth.token_ttl"))
	authhttp.RegisterHTTPEndpoints(router, app.authUC)
	app.httpServer = &http.Server{
		Addr:         appCfg.Address,
		Handler:      router,
		ReadTimeout:  appCfg.ReadTimeout,
		WriteTimeout: appCfg.WriteTimeout,
		IdleTimeout:  appCfg.IdleTimeout,
	}
	return app
}

func (app *App) Run() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Printf("⇨ http server started on %s\n", app.httpServer.Addr)
		log.Fatal(app.httpServer.ListenAndServe())
	}()
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return app.httpServer.Shutdown(ctx)
}
