package main

import (
	"account/dao/mysql"
	"account/dao/redis"
	"account/routers"
	"log"
)

// @title 记账后端接口文档
// @version 1.0
// @description 记账操作
// @termsOfService http://swagger.io/terms/

// @contact.name 青烟绕指柔
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath
func main() {
	err := mysql.InitMysql()
	if err != nil {
		log.Println("Run mysql failed : ", err.Error())
		return
	}
	log.Println("---------- Run mysql success ! ----------")
	defer mysql.MysqlClose()

	err = redis.InitRedis()
	if err != nil {
		log.Println("Run redis failed : ", err.Error())
		return
	}
	log.Println("---------- Run redis success ! ----------")
	defer redis.RedisClose()

	r := routers.NewAPI()
	if err = r.Run(":8080"); err != nil {
		log.Println(err.Error())
	}
}
