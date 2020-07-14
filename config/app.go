package config

import (
"fmt"
"github.com/spf13/viper"
"sync"
"time"
)

// App hold the configuration for the main app
type App struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Debug        bool
	AppKeys      []string
	IpnURL       string
}

var app *App

// loadApp loads the main app's configuration from consul with viper
func loadApp() {
	app = &App{
		Address:      fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port")),
		ReadTimeout:  viper.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration("app.write_timeout") * time.Second,
		IdleTimeout:  viper.GetDuration("app.idle_timeout") * time.Second,
		Debug:        viper.GetBool("app.debug"),
		AppKeys:      viper.GetStringSlice("app.app_keys"),
		IpnURL:       viper.GetString("app.ipn_url"),
	}
}

// GetApp returns the App instance
func GetApp() *App {
	var loadAppOnce sync.Once
	loadAppOnce.Do(loadApp)
	return app
}
