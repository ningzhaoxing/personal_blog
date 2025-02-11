package initialize

import (
	"fmt"
	"gin-blog/pkg/globals"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDb(c *globals.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DB.Config.User, c.DB.Config.Password, c.DB.Config.Host, c.DB.Config.Port, c.DB.Config.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Sprint("Mysql connect err:", err.Error()))
		return nil
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Sprint("sql db error：", err.Error()))
		return nil
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDb.SetMaxIdleConns(c.DB.Config.MaxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(c.DB.Config.MaxCon)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)

	return db
}
