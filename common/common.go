package common

import (
	"github.com/jinzhu/gorm"
)

// User 用户信息
type User struct {
	UserName    string `gorm:"size:2048" json:"username" form:"username"` // 登录账户
	Pwd         string `gorm:"size:2048" json:"password" form:"password"` // 登录密码
	Name        string `gorm:"size:2048" json:"name" form:"name"`         // 用户姓名
	PhoneNumber string `gorm:"size:2048" json:"phone" form:"phone"`       // 电话号码
	IsHide      bool   `json:"is_hide" form:"is_hide"`   // 是否隐藏自己的信息
	Gender      int    `json:"gender" form:"gender"`     // 性别
	Birthday    string `gorm:"size:2048" json:"birthday" form:"birthday"` //生日
	Solar       int    `json:"solar" form:"solar"`       //
	gorm.Model
}

// Image 图片信息
type Image struct {
	Name     string `json:"name" form:"name"`         // 图片名字
	UserName string `json:"username" form:"username"` // 图片上传者
	gorm.Model
}

// Account 记账条目信息
type Account struct {
	Action   bool    `json:"action" form:"action"`     // 支出 true 收入 false
	Name     string  `json:"name" form:"name"`         // 账单条目名字
	Type     int64   `json:"type" form:"type"`         // 账单条目类型 衣、食、住、行
	Desc     string  `json:"desc" form:"desc"`         // 账单条目描述
	UserName string  `json:"username" form:"username"` // 账单上传者
	Money    float32 `json:"money" form:"money"`       // 费用
	Extra    string  `json:"extra" from:"extra"`       // 额外字段，以json格式存储
	gorm.Model
}

const (
	Dress int64 = iota
	Eat
	Stay
	Travel
)

var AccountDesc = map[string]int{
	"Dress":  0,
	"Eat":    1,
	"Stay":   2,
	"Travel": 3,
}

const Host = "127.0.0.1:8080"
