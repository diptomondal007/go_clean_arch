package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"time"
)

// DB holds the configuration for the default db
type DB struct {
	Address         string
	User            string
	Name            string
	Password        string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
}

var db *DB

// loadDB load the default db's configuration from consul with viper
func loadDB() {
	db = &DB{
		Address:         fmt.Sprintf("%s:%d", viper.GetString("database.host"), viper.GetInt("database.port")),
		User:            viper.GetString("database.user"),
		Name:            viper.GetString("database.name"),
		Password:        viper.GetString("database.password"),
		MaxIdleConn:     viper.GetInt("database.max_idle_connections"),
		MaxOpenConn:     viper.GetInt("database.max_open_connections"),
		MaxConnLifetime: viper.GetDuration("database.max_connection_lifetime") * time.Second,
		Debug:           viper.GetBool("database.debug"),
	}
}


// GetDB returns the App instance
func GetDB() *DB {
	var loadDBOnce sync.Once
	loadDBOnce.Do(loadDB)
	return db
}
