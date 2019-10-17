package godb

import (
	"github.com/ikaiguang/go-utils/db/config"
	"github.com/ikaiguang/go-utils/db/mysql"
	"github.com/ikaiguang/go-utils/db/postgres"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strings"
)

// tablePrefix db table prefix
var tablePrefix string

// TablePrefix table prefix
func TablePrefix() string {
	return tablePrefix
}

// db
var (
	dbPool *gorm.DB // db conn
)

// database driver
const (
	DbDriverMysql    = "mysql"    // mysql
	DbDriverPostgres = "postgres" // postgres
)

// GetDBPool db pool
func GetDBPool() (*gorm.DB, error) {
	if dbPool != nil {
		return dbPool, nil
	}
	dbConn, err := NewDBConn()
	if err != nil {
		return dbConn, errors.WithStack(err)
	}
	dbPool = dbConn
	return dbPool, nil
}

// GetDBPoolWithConfig db pool
func GetDBPoolWithConfig(cfg *configs.Config) (*gorm.DB, error) {
	if dbPool != nil {
		return dbPool, nil
	}
	dbConn, err := NewDBConnWithConfig(cfg)
	if err != nil {
		return dbConn, errors.WithStack(err)
	}
	dbPool = dbConn
	return dbPool, nil
}

// NewDBConn : db conn
var NewDBConn = func() (*gorm.DB, error) {
	return newDBConn(configs.InitConfig())
}

// NewDBConnWithConfig db conn
func NewDBConnWithConfig(cfg *configs.Config) (*gorm.DB, error) {
	return newDBConn(cfg)
}

// newDBConn : db conn
func newDBConn(cfg *configs.Config) (*gorm.DB, error) {
	var dbConn *gorm.DB
	var err error

	// open db connection
	switch cfg.Driver {
	case DbDriverMysql: // mysql
		dbConn, err = mysql.NewDBConn(cfg)

	case DbDriverPostgres: // postgres
		dbConn, err = postgres.NewDBConn(cfg)

	default: // invalid database driver
		return nil, errors.New("invalid database driver")
	}

	// db connection error
	if err != nil {
		return dbConn, errors.WithStack(err)
	}

	// table prefix
	if cfg.SetTablePrefix {
		SetTablePrefix(cfg)
	}

	// debug
	dbConn = SetDebug(dbConn, cfg)

	// max open conn
	SetMaxOpenConn(dbConn, cfg)

	// max idle conn
	SetMaxIdleConn(dbConn, cfg)

	// conn max lifetime
	SetConnMaxLifetime(dbConn, cfg)

	// other option
	dbConn, err = SetConnOptions(dbConn)
	if err != nil {
		return dbConn, errors.WithStack(err)
	}

	// ping
	if err := dbConn.DB().Ping(); err != nil {
		return dbConn, errors.WithStack(err)
	}
	return dbConn, nil
}

// SetConnOptions 设置选项
var SetConnOptions = func(dbConn *gorm.DB) (*gorm.DB, error) { return dbConn, nil }

// SetConnMaxLifetime : conn max lifetime
// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are reused forever.
func SetConnMaxLifetime(dbConn *gorm.DB, cfg *configs.Config) {
	if cfg.MaxLifetime <= 0 {
		return
	}
	dbConn.DB().SetConnMaxLifetime(cfg.MaxLifetime)
}

// SetMaxIdleConn : set max idle conn
// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in
// a future release.
func SetMaxIdleConn(dbConn *gorm.DB, cfg *configs.Config) {
	if cfg.MaxIdle <= 0 {
		return
	}
	dbConn.DB().SetMaxIdleConns(cfg.MaxIdle)
}

// SetMaxOpenConn : set max open conn
// SetMaxOpenConns sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func SetMaxOpenConn(dbConn *gorm.DB, cfg *configs.Config) {
	if cfg.MaxOpen <= 0 {
		return
	}
	dbConn.DB().SetMaxOpenConns(cfg.MaxOpen)
}

// SetDebug : print debug
// LogMode set log mode, `true` for detailed logs, `false` for no log,
// default, will only print error logs
func SetDebug(dbConn *gorm.DB, cfg *configs.Config) *gorm.DB {
	return dbConn.LogMode(cfg.Debug)
}

// SetTablePrefix : set table prefix
var SetTablePrefix = func(cfg *configs.Config) {
	tablePrefix = cfg.TablePrefix
	// rewrite handler
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		if !strings.HasPrefix(tableName, tablePrefix) {
			return tablePrefix + tableName
		}
		return tableName
	}
}
