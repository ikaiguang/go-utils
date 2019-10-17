package godb

import (
	"github.com/ikaiguang/go-utils/db/config"
	"os"
	"testing"
)

func TestNewDBConn(t *testing.T) {
	// option
	os.Setenv(configs.EnvKeyDbDebug, "true")
	os.Setenv(configs.EnvKeyDbMaxOpen, "10")
	os.Setenv(configs.EnvKeyDbMaxIdle, "10")
	os.Setenv(configs.EnvKeyDbMaxLifetime, "30s")

	// mysql
	t.Log("test mysql ... \n")
	// auth
	// root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local
	os.Setenv(configs.EnvKeyDbDriver, DbDriverMysql)
	os.Setenv(configs.EnvKeyDbAddress, "127.0.0.1:3306")
	os.Setenv(configs.EnvKeyDbName, "test")
	os.Setenv(configs.EnvKeyDbTablePrefix, "ag_")
	os.Setenv(configs.EnvKeyDbUser, "root")
	os.Setenv(configs.EnvKeyDbPassword, "Mysql.123456")
	os.Setenv(configs.EnvKeyDbParameters, "charset=utf8&timeout=60s&loc=Local&autocommit=true")
	db, err := NewDBConn()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%v \n", db)

	// postgres
	t.Log("test postgres ... \n")
	// auth
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword
	// postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
	os.Setenv(configs.EnvKeyDbDriver, DbDriverPostgres)
	os.Setenv(configs.EnvKeyDbAddress, "127.0.0.1:5432")
	os.Setenv(configs.EnvKeyDbName, "postgres")
	os.Setenv(configs.EnvKeyDbTablePrefix, "ag_")
	os.Setenv(configs.EnvKeyDbUser, "postgres")
	os.Setenv(configs.EnvKeyDbPassword, "Postgres.123456")
	os.Setenv(configs.EnvKeyDbParameters, "connect_timeout=60&sslmode=disable")
	db, err = NewDBConn()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%v \n", db)
}
