package mysql

import (
	. "account/common"
	"account/dao/mysql"
	"log"
)

func GetUser(name, pwd string) bool {
	res := 0
	if err := mysql.MysqlDB.Model(&User{}).Where("user_name = ? AND pwd = ?", name, pwd).Count(&res); err != nil {
		return res > 0
	}
	return false
}

func CountUserByUserName(userName string) bool {
	res := 0
	if err := mysql.MysqlDB.Model(&User{}).Where("user_name = ?", userName).Count(&res); err != nil {
		return res > 0
	}
	return false
}

func AddUser(user *User) error {
	log.Println("In AddUser")
	log.Println(user)
	return mysql.MysqlDB.Create(&user).Error
}

func GetUserByID(id int64) bool {
	res := 0
	if err := mysql.MysqlDB.Model(&User{}).Where("id = ?", id).Count(&res); err != nil {
		return res > 0
	}
	return false
}

func GetUserByUserName(userName string) (res *User, err error) {
	res = new(User)
	err = mysql.MysqlDB.Model(&User{}).Where("user_name = ?", userName).Find(&res).Error
	return res, err
}

func GetUserByIDs(id []int64) (res []*User, err error) {
	log.Println(id)
	err = mysql.MysqlDB.Model(&User{}).Where("id IN (?)", id).Find(&res).Error
	return res, err
}
