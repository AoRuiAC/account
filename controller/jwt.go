package controller

import (
	"account/common"
	"account/models/mysql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

type MyClaims struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	jwt.StandardClaims
}

var accessSecret = []byte("qyrzr")

// GetToken 获取accessToken和refreshToken
func GetToken(name, pwd string) string {
	// accessToken 的数据
	aT := MyClaims{
		name,
		pwd,
		jwt.StandardClaims{
			Issuer:    "AR",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * 30 * time.Hour).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aT)
	accessTokenSigned, err := accessToken.SignedString(accessSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return ""
	}
	return accessTokenSigned
}

func ParseToken(accessTokenString string) (*MyClaims, error) {
	fmt.Println("In ParseToken")
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if claims, ok := accessToken.Claims.(*MyClaims); ok && accessToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 用鉴权到中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 默认双Token放在请求头Authorization的Bearer中，并以空格隔开
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(c.Request.Header)
		if authHeader == "" {
			c.JSON(200, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		fmt.Println("authHeader = ", authHeader)
		parts := strings.Split(authHeader, " ")
		fmt.Println("len = ", len(parts))
		fmt.Println("parts[0] = ", parts[0])
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			c.JSON(200, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		parseToken, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(200, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		parts[1] = GetToken(parseToken.Name, parseToken.Pwd)
		// 如果需要刷新双Token时，返回双Token
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "鉴权成功",
			"data": gin.H{
				"token": parts[1],
			},
		})

		c.Set("username", parseToken.Name)
		c.Next()
	}
}

// @Summary 获取token接口
// @Description 使用jwt鉴权，默认双Token放在请求头Authorization的Bearer中，并以空格隔开
// @Tags jwt
// @Accept application/json
// @Produce application/json
// @Param user formData string true "用户名" default(admin)
// @Success 200 "{"msg": "success"}"
// @Failure 300 "{"msg": "无效参数"}"
// @Failure 301 "{"msg": "鉴权失败，用户不存在"}"
// @Router /link [delete]
func AuthHandler(c *gin.Context) {
	log.Println("In authHandler")
	var user common.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(300, gin.H{
			"msg": "无效参数",
		})
		log.Println(err.Error())
		return
	}
	log.Println("user = ", user)
	if !mysql.GetUser(user.UserName, user.Pwd) {
		c.JSON(301, gin.H{
			"msg": "鉴权失败，用户不存在",
		})
		log.Println("User not exist or password error")
		return
	}

	userMsg, err := mysql.GetUserByUserName(user.UserName)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "后端出错",
		})
		return
	}

	accessTokenString := GetToken(user.UserName, user.Pwd)
	c.JSON(200, gin.H{
		"msg":     "success",
		"success": true,
		"data": gin.H{
			"token": accessTokenString,
			"user": gin.H{
				"username": userMsg.UserName,
				"password": userMsg.Pwd,
				"phone":    userMsg.PhoneNumber,
				"name":     userMsg.Name,
				"gender":   userMsg.Gender,
				"birthday": userMsg.Birthday,
				"solar":    userMsg.Solar,
			},
		},
	})
}
