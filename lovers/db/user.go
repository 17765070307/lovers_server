package db

import (
	"time"

	"github.com/gin-gonic/gin"
)

//
//  UserInfo
//  @Description: 用户信息
//
type User struct {
	Id        int64     `gorm:"type:int;not null;column:id;comment:主键;autoIncrement" json:"id"`
	Code      string    `gorm:"type:varchar(32);not null;column:user_code;comment:用户编号" json:"user_code"`
	Password  string    `gorm:"type:varchar(32);null;column:user_password;comment:用户密码" json:"user_password"`
	Name      string    `gorm:"type:varchar(32);null;column:user_name;comment:用户姓名" json:"user_name"`
	Sex       string    `gorm:"type:varchar(32);null;column:user_sex;comment:用户性别" json:"user_sex"`
	Weight    float32   `gorm:"type:double;null;column:user_weight;default:0;comment:用户当前体重" json:"user_weight"`
	Phone     string    `gorm:"type:varchar(32);null;column:user_phone;comment:用户手机号" json:"user_phone"`
	CreatedAt time.Time `gorm:"type:datetime;null;column:create_time;comment:用户创建时间" json:"create_time"`
	UpdatedAt time.Time `gorm:"type:datetime;null;column:modify_time;comment:更新时间" json:"modify_time"`
}

func (User) TableName() string {
	return "user"
}

//
//  UserQuery
//  @Description: 查询用户信息
//  @param ctx 上下文信息
//  @param sql sql语句
//  @return []User 用户信息列表
//  @return error 异常信息
//
func UserQuery(ctx *gin.Context) ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return []User{}, result.Error
	}
	return users, nil
}
