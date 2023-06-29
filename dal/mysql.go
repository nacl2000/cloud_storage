package dal

import (
	"fmt"
	"github.com/nacl2000/clould_storage/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBHandler *gorm.DB

func InitDB() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		conf.DBInfo.User, conf.DBInfo.Passwd, conf.DBInfo.Host,
		conf.DBInfo.Port, conf.DBInfo.DBName)
	var err error
	DBHandler, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDBHandler() *gorm.DB {
	return DBHandler
}

func GetModelDB(model interface{}) *gorm.DB {
	return DBHandler.Model(model)
}
