package controller

import (
	"account/models/mysql"
	"account/models/redis"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// @Summary 用户关系建立接口
// @Description 用户进行关系建立接口
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "设置成功"}"
// @Failure 300 "{"msg": "上传的id无效，请重新上传"}"
// @Failure 301 "{"msg": "需要绑定的用户不存在"}"
// @Failure 500 "{"msg": "服务器抖动，请稍后重试"}"
// @Router /link/:id1/:id2 [post]
func SetLinkHandler(c *gin.Context) {
	log.Println("In SetLinkHandler")
	id1, ok := c.Params.Get("id1")
	if !ok {
		c.JSON(300, gin.H{
			"msg":  "上传的id无效，请重新上传",
		})
		return
	}
	id2, ok := c.Params.Get("id2")
	if !ok {
		c.JSON(300, gin.H{
			"msg":  "上传的id无效，请重新上传",
		})
		return
	}
	ID1, _ := strconv.ParseInt(id1, 10, 64)
	ID2, _ := strconv.ParseInt(id2, 10, 64)
	if !mysql.GetUserByID(ID1) || !mysql.GetUserByID(ID2) {
		c.JSON(301, gin.H{
			"msg":  "需要绑定的用户不存在",
		})
		return
	}
	err := redis.SetLink(id1, id2)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"msg":  "服务器抖动，请稍后重试",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "设置成功",
	})
}

// @Summary 获取用户绑定的用户接口
// @Description 获取用户绑定的用户
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "设置成功", "data": "res"}"
// @Failure 300 "{"msg": "上传的id无效，请重新上传"}"
// @Failure 301 "{"msg": "需要绑定的用户不存在"}"
// @Failure 500 "{"msg": "服务器抖动，请稍后重试"}"
// @Router /link/:id [get]
func GetLinkHandler(c *gin.Context) {
	log.Println("In GetLinkHandler")
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(300, gin.H{
			"msg":  "上传的id无效，请重新上传",
		})
		return
	}
	ID, _ := strconv.ParseInt(id, 10, 64)
	if !mysql.GetUserByID(ID) {
		c.JSON(301, gin.H{
			"msg":  "需要获取的用户不存在",
		})
		return
	}
	res, err := redis.GetLinkById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "服务器抖动，请稍后重试",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "获取成功",
		"data": res,
	})
}

// @Summary 删除用户绑定接口
// @Description 删除用户1和用户2之间的关系
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "删除成功"}"
// @Failure 300 "{"msg": "上传的id无效，请重新上传"}"
// @Failure 301 "{"msg": "删除绑定的用户不存在"}"
// @Failure 500 "{"msg": "服务器抖动，请稍后重试"}"
// @Router /link/:id1/:id2 [delete]
func DeleteLinkHandler(c *gin.Context) {
	log.Println("In DeleteLinkHandler")
	id1, ok := c.Params.Get("id1")
	if !ok {
		c.JSON(300, gin.H{
			"msg":  "上传的id无效，请重新上传",
		})
		return
	}
	id2, ok := c.Params.Get("id2")
	if !ok {
		c.JSON(300, gin.H{
			"msg":  "上传的id无效，请重新上传",
		})
		return
	}
	ID1, _ := strconv.ParseInt(id1, 10, 64)
	ID2, _ := strconv.ParseInt(id2, 10, 64)
	if !mysql.GetUserByID(ID1) || !mysql.GetUserByID(ID2) {
		c.JSON(301, gin.H{
			"msg":  "删除绑定的用户不存在",
		})
		return
	}
	err := redis.DeleteLink(id1, id2)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "服务器抖动，请稍后重试",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "删除成功",
	})
}
