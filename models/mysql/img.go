package mysql

import (
	. "account/common"
	"account/dao/mysql"
	"log"
)

func AddImg(img *Image) error {
	return mysql.MysqlDB.Create(&img).Error
}

func DeleteImgById(id int64) error {
	return mysql.MysqlDB.Delete(&Image{}).Where("id = ?", id).Error
}

func DeleteImgByName(name string) error {
	log.Println("In DeleteImgByName")
	log.Println("name = ", name)
	return mysql.MysqlDB.Where("name = ?", name).Delete(&Image{}).Error
}

func GetImgByName(userName string) (res []*Image, err error) {
	err = mysql.MysqlDB.Model(&Image{}).Where("user_name = ?", userName).Find(&res).Error
	return res, err
}

func GetImgByNames(userNames []string) (res []*Image, err error) {
	err = mysql.MysqlDB.Model(&Image{}).Where("user_name IN (?)", userNames).Find(&res).Error
	return res, err
}
