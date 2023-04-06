package controller

import (
	. "account/common"
	"account/models/mysql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// @Summary 添加账单条目
// @Description 添加账单条目
// @Tags account
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "添加成功"}"
// @Failure 301 "{"msg": "请重新上传信息"}"
// @Failure 400 "{"msg": "账单条目类型错误"}"
// @Failure 500 "{"msg": "服务器出错，请稍后重试"}"
// @Router /account [post]
func PostAccountHandler(c *gin.Context) {
	var account Account
	err := c.ShouldBind(&account)
	if err != nil {
		c.JSON(301, gin.H{
			"msg": "请重新上传信息",
		})
		log.Println(err.Error())
		return
	}

	if account.Type < 0 || account.Type > 3 {
		c.JSON(400, gin.H{
			"msg": "账单条目类型错误",
		})
		return
	}

	err = mysql.AddAccount(&account)
	if err != nil {
		log.Println("AddAccount err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "添加成功",
	})
}

// @Summary 获取用户账单条目
// @Description 获取用户账单条目
// @Tags account
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "获取成功","data": "res"}"
// @Failure 500 "{"msg": "服务器出错，请稍后重试"}"
// @Router /account [get]
func GetAccountHandler(c *gin.Context) {
	userName := c.MustGet("username").(string)
	accounts, err := mysql.GetAccountByUserName(userName)
	if err != nil {
		log.Println("GetAccountByUserName err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}

	res, err := json.Marshal(accounts)

	if err != nil {
		log.Println("Marshal err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "获取成功",
		"data": string(res),
	})
}

// @Summary 删除用户账单条目
// @Description 按照账单id删除用户账单条目
// @Tags account
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "删除成功"}"
// @Failure 300 "{"msg": "上传的id无效，请重新上传"}"
// @Failure 500 "{"msg": "服务器出错，请稍后重试"}"
// @Router /account/:id [delete]
func DeleteAccountHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(300, gin.H{
			"msg": "上传的id无效，请重新上传",
		})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(300, gin.H{
			"msg": "上传的id无效，请重新上传",
		})
		return
	}
	err = mysql.DeleteAccountByID(int64(idInt))
	if err != nil {
		log.Println("DeleteAccountByID err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}
