package database

import (
	"fmt"
	"github.com/Niexiawei/golang-skeleton/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	_conn       *gorm.DB
	mysqlConfig config.Mysql
)

func GetDatabase() *gorm.DB {
	return _conn
}

func SetupDatabase() {
	var err error
	mysqlConfig = config.Instance.Mysql
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Db,
	)

	_conn, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: func() logger.Interface {
			return logger.Default.LogMode(logger.Warn)
		}(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	conn, err := _conn.DB()
	if err != nil {
		panic(err)
	}
	conn.SetMaxOpenConns(mysqlConfig.Pool)
	conn.SetMaxIdleConns(mysqlConfig.Pool)
}
