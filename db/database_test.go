package godb

import (
	"github.com/ikaiguang/go-utils/db/config"
	"os"
	"testing"
)

func TestNewDBConn(t *testing.T) {
	// option
	os.Setenv(godbconfigs.EnvKeyDbDebug, "true")
	os.Setenv(godbconfigs.EnvKeyDbMaxOpen, "10")
	os.Setenv(godbconfigs.EnvKeyDbMaxIdle, "10")
	os.Setenv(godbconfigs.EnvKeyDbMaxLifetime, "30s")

	// mysql
	t.Log("test mysql ... \n")
	// auth
	// root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local
	os.Setenv(godbconfigs.EnvKeyDbDriver, DbDriverMysql)
	os.Setenv(godbconfigs.EnvKeyDbEndpoints, "127.0.0.1:3306")
	os.Setenv(godbconfigs.EnvKeyDbName, "test")
	os.Setenv(godbconfigs.EnvKeyDbTablePrefix, "ag_")
	os.Setenv(godbconfigs.EnvKeyDbUser, "root")
	os.Setenv(godbconfigs.EnvKeyDbPassword, "Mysql.123456")
	os.Setenv(godbconfigs.EnvKeyDbParameters, "charset=utf8&timeout=60s&loc=Local&autocommit=true")
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
	os.Setenv(godbconfigs.EnvKeyDbDriver, DbDriverPostgres)
	os.Setenv(godbconfigs.EnvKeyDbEndpoints, "127.0.0.1:5432")
	os.Setenv(godbconfigs.EnvKeyDbName, "postgres")
	os.Setenv(godbconfigs.EnvKeyDbTablePrefix, "ag_")
	os.Setenv(godbconfigs.EnvKeyDbUser, "postgres")
	os.Setenv(godbconfigs.EnvKeyDbPassword, "Postgres.123456")
	os.Setenv(godbconfigs.EnvKeyDbParameters, "connect_timeout=60&sslmode=disable")
	db, err = NewDBConn()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%v \n", db)
}
