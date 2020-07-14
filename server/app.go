package server

import (
	"context"
	"github.com/diptomondal007/go_clean_arch/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp(appCfg *config.App) *App {
	router := echo.New()

	// middleware
	router.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	return &App{
		httpServer: &http.Server{
			Addr:         appCfg.Address,
			Handler:      router,
			ReadTimeout:  appCfg.ReadTimeout,
			WriteTimeout: appCfg.WriteTimeout,
			IdleTimeout:  appCfg.IdleTimeout,
		},
	}
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
