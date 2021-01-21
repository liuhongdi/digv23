package global

import (
	"gorm.io/gorm"
	"time"
	"gorm.io/driver/mysql"
)

var (
	DBLink *gorm.DB
)

//创建mysql链接
func SetupDBLink() (error) {
	var err error
	dsn:="root:password@tcp(127.0.0.1:3306)/business?charset=utf8&parseTime=True&loc=Local";
	DBLink, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DBLink.Logger.LogMode(true)
	if err != nil {
		return err
	}
	sqlDB, err := DBLink.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(30)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
