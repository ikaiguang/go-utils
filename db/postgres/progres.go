package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" // driver
	"strings"

	"github.com/ikaiguang/go-utils/db/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// NewDBConn : db conn
func NewDBConn(cfg *configs.Config) (*gorm.DB, error) {
	// db connection
	dbConn, err := gorm.Open("postgres", InitPostgresDsn(cfg))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dbConn, err
}

// InitPostgresDsn : dsn = "postgresql://postgres:Postgres.123456@127.0.0.1:5432/postgres?connect_timeout=60&sslmode=disable"
//
// dsn layout postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
// or
// dsn layout host=myhost port=myport user=postgres dbname=postgres password=mypassword
//
// https://www.postgresql.org/docs/11/libpq-connect.html#LIBPQ-CONNSTRING
var InitPostgresDsn = func(cfg *configs.Config) string {
	var dsn = "postgresql://"

	// user
	if len(cfg.Username) > 0 {
		dsn += cfg.Username
		if len(cfg.Password) > 0 {
			dsn += ":" + cfg.Password
		}
		dsn += "@"
	}

	// address
	dsn += strings.Join(cfg.Address, ",")

	// name
	if len(cfg.DBName) > 0 {
		dsn += "/" + cfg.DBName
	}

	// parameters
	if len(cfg.Parameters) > 0 {
		dsn += "?" + cfg.Parameters
	}
	return dsn
}
