package controller

import (
	"account/common"
	"account/models/mysql"
	"account/models/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
)

// @Summary 获取用户以及用户关注者上传的图片
// @Description 获取用户以及用户关注者上传的图片
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "success"}"
// @Failure 500 "{"msg": "服务器出错，请稍后重试"}"
// @Router /img [delete]
func GetImageHandler(c *gin.Context) {
	log.Println("In GetImageHandler")
	username := c.MustGet("username").(string)

	user, err := mysql.GetUserByUserName(username)
	if err != nil {
		log.Println("GetUserByUserName err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}

	users, err := redis.GetLinkById(strconv.Itoa(int(user.ID)))
	if err != nil {
		log.Println("GetUserByName err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}

	newUsers := make([]int64, 0)
	for _, user := range users {
		userInt, _ := strconv.ParseInt(user, 10, 64)
		newUsers = append(newUsers, userInt)
	}

	log.Println("newUsers = ", newUsers)

	AllUser, err := mysql.GetUserByIDs(newUsers)
	if err != nil {
		log.Println("GetUserByIDs err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}

	AllName := make([]string, 0)
	for _, user := range AllUser {
		if user.IsHide == true { // 被隐藏就不获取
			continue
		}
		AllName = append(AllName, user.UserName)
	}
	AllName = append(AllName, user.UserName) // 加入本人的图片

	fmt.Println("AllName = ", AllName)

	Imgs, err := mysql.GetImgByNames(AllName)
	if err != nil {
		log.Println("GetImgByNames err : ", err)
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		return
	}

	fmt.Println("Imgs = ", Imgs)

	files := make([]string, 0)
	for _, img := range Imgs {
		file := common.Host + "/static/" + img.Name
		files = append(files, file)
	}

	c.JSON(200, gin.H{
		"msg": files,
	})
}

// @Summary 添加图片
// @Description 添加图片
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "success"}"
// @Failure 301 "{"msg": "请使用jpg或者png格式的图片上传"}"
// @Failure 500 "{"msg": "服务器出错，请稍后重试"}"
// @Failure 501 "{"msg": "保存图片出错，请稍后重试"}"
// @Router /img [post]
func PostImageHandler(c *gin.Context) {
	log.Println("In PostImageHandler")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "服务器出错，请稍后重试",
		})
		log.Println(err.Error())
		return
	}
	ext := strings.Split(file.Filename, ".")
	if ext[1] != "jpg" && ext[1] != "png" {
		c.JSON(301, gin.H{
			"msg": "请使用jpg或者png格式的图片上传",
		})
		return
	}
	filePath := "./img/"
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(filePath, os.ModePerm)
			if err != nil {
				log.Println("[PostImageHandler] Mkdir error : ", err)
				return
			}
		}
	}
	filePath = filePath + file.Filename
	err = c.SaveUploadedFile(file, filePath) // 把文件存在本地
	if err != nil {
		c.JSON(501, gin.H{
			"msg": "保存图片出错，请稍后重试",
		})
		log.Println("saveUploadedFile error : ", err.Error())
		return
	}

	err = mysql.AddImg(&common.Image{
		Name:     file.Filename,
		UserName: c.MustGet("username").(string),
	})
	if err != nil {
		c.JSON(501, gin.H{
			"msg": "保存图片出错，请稍后重试",
		})
		log.Println("AddImg error : ", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"msg": "保存图片成功",
	})
}

// @Summary 删除图片
// @Description 根据图片名字删除图片
// @Tags image
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "删除成功"}"
// @Failure 301 "{"msg": "删除的图片名字无效，请重新上传"}"
// @Failure 500 "{"msg": "服务器出错，请稍后重试"}"
// @Router /img/:name [post]
func DeleteImageHandler(c *gin.Context) {
	log.Println("In DeleteImageHandler")
	name, ok := c.Params.Get("name")
	log.Println("name = ", name)
	if !ok {
		c.JSON(301, gin.H{
			"msg": "删除的图片名字无效，请重新上传",
		})
	}

	err := os.RemoveAll("./img/" + name)
	if err != nil {
		log.Println("RemoveAll err : ", err)
		c.JSON(300, gin.H{
			"msg": err,
		})
		return
	}

	err = mysql.DeleteImgByName(name)
	if err != nil {
		log.Println("DeleteImgByName err : ", err)
		c.JSON(500, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}
