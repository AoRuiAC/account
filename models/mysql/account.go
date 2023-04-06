package mysql

import (
	. "account/common"
	"account/dao/mysql"
)

func AddAccount(account *Account) error {
	return mysql.MysqlDB.Create(&account).Error
}

func GetAccountByUserName(userName string) (account []*Account, err error) {
	err = mysql.MysqlDB.Model(&Account{}).Where("user_name = ?", userName).Find(&account).Error
	return account, err
}

func DeleteAccountByID(id int64) error {
	return mysql.MysqlDB.Where("id = ?", id).Delete(&Account{}).Error
}
