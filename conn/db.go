package conn

import (
	"fmt"
	"github.com/diptomondal007/go_clean_arch/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

const (
	dialect = "mysql"
)

// DB holds the database instance
type DB struct {
	*gorm.DB
}

//defaultDB is the default database instance
var defaultDB *DB

//connectDB connect and set db client of database using db configuration
func connectDB(cfg *config.DB) error{
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", cfg.User, cfg.Password, cfg.Address, cfg.Name)
	conn, err := gorm.Open(dialect, uri)
	if err != nil{
		return err
	}
	defaultDB = &DB{conn}
	defaultDB.DB.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	defaultDB.DB.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	defaultDB.DB.DB().SetConnMaxLifetime(cfg.MaxConnLifetime)
	defaultDB.DB.LogMode(cfg.Debug)

	return nil
}

func GetDB() (*DB, error){
	var err error
	var connDBOnce sync.Once
	connDBOnce.Do(func() {
		err = connectDB(config.GetDB())
	})
	return defaultDB, err
}