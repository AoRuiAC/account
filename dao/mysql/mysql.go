package mysql

import (
	. "account/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDB *gorm.DB

func InitMysql() (err error) {
	MysqlDB, err = gorm.Open("mysql", "root:123456@(124.221.179.105)/account?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	MysqlDB.AutoMigrate(&User{}, &Image{}, &Account{})
	return MysqlDB.DB().Ping()
}

func MysqlClose() {
	if err := MysqlDB.Close(); err != nil {
		panic(err)
	}
}
