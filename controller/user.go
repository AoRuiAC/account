package controller

import (
	"account/common"
	"account/models/mysql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 用户注册接口
// @Description 用户进行注册的接口
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "上传成功"}"
// @Failure 300 "{"msg": "请重新上传信息"}"
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	log.Println("In RegisterHandler")
	var user common.User
	log.Println("register user:", c.Request)
	log.Println(c.Request.Header)
	log.Println(c.Request.Body)

	jsonParam := c.Query("json")
	log.Println(jsonParam)

	err := json.Unmarshal([]byte(jsonParam), &user)

	// err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(300, gin.H{
			"msg": "请重新上传信息",
		})
		log.Println(err.Error())
		return
	}
	log.Println("user = ", user)
	if user.UserName == "" || user.Name == "" || user.Pwd == "" {
		c.JSON(300, gin.H{
			"msg": "未填写必填信息，请重新上传信息",
		})
		log.Println("未填写必填信息，请重新上传信息")
		return
	}

	if mysql.CountUserByUserName(user.UserName) {
		c.JSON(300, gin.H{
			"msg": "用户已存在",
		})
		return
	}

	err = mysql.AddUser(&user)
	log.Println("AddUser error = ", err)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		log.Println("err = ", err.Error())
		return
	}

	token := GetToken(user.UserName, user.Pwd)

	c.JSON(200, gin.H{
		"msg":     "上传成功",
		"success": true,
		"data": gin.H{
			"user": gin.H{
				"username": user.UserName,
				"password": user.Pwd,
				"phone":    user.PhoneNumber,
				"name":     user.Name,
				"gender":   user.Gender,
				"birthday": user.Birthday,
				"solar":    user.Solar,
			},
			"token": token,
		},
	})
}
