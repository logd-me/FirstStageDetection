/*
@Title : User
@Description :
@Author : 谭靖渝
@Update : 2021/5/4 17:25
*/
package Model

import (
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// User 用户模型
type User struct {
	CreatedAt time.Time `gorm:"comment:'创建时间';type:datetime;"`
	UpdatedAt time.Time `gorm:"comment:'修改时间';type:datetime;"`
	UID       string    `json:"uid" form:"uid" gorm:"primary_key"`
	PassWord  string    `json:"password" form:"password" gorm:"column:password"`
	NickName  string    `json:"nickname" form:"nickname" gorm:"column:nickname"`
	Status    string    `json:"status" form:"status"`
	Avatar    string    `gorm:"size:1000"`
	Article   []Article
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

//注册账号
func Random(width int) string {
	var rands = 0
	for i := 0; i < width; i++ {
		r := rand.New(rand.NewSource(time.Now().Add(time.Second * time.Duration(i)).UnixNano()))
		n := r.Intn(10)
		rands = rands + n*int(math.Pow10(width-i-1))
	}
	str := strconv.Itoa(rands)
	return str
}

func GetUser(UID interface{}) (User, error) {
	var user User
	result := DB.First(&user, UID)
	return user, result.Error
}

//设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
}

//校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return err == nil
}

