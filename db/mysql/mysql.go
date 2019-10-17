package mysql

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" // driver
	"strings"

	"github.com/ikaiguang/go-utils/db/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// NewDBConn : db conn
func NewDBConn(cfg *configs.Config) (*gorm.DB, error) {
	// db connection
	dbConn, err := gorm.Open("mysql", InitMysqlDsn(cfg))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dbConn, err
}

// InitMysqlDsn : dsn = "root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&timeout=60s&loc=Local&autocommit=true"
//
// dsn layout [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
//
// github.com/go-sql-driver/mysql -> mysql.Config{}.FormatDSN()
var InitMysqlDsn = func(cfg *configs.Config) string {
	var dsn string

	// user
	if len(cfg.Username) > 0 {
		dsn += cfg.Username
		if len(cfg.Password) > 0 {
			dsn += ":" + cfg.Password
		}
		dsn += "@"
	}

	// address
	if len(cfg.Address) > 0 {
		dsn += "tcp(" + strings.Join(cfg.Address, ",") + ")"
	}

	// name
	dsn += "/" + cfg.DBName

	// parameters
	if len(cfg.Parameters) > 0 {
		dsn += "?" + cfg.Parameters
	}
	return dsn
}
