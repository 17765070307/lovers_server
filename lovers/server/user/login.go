package user

import (
	"Lovers/db"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, code, password string) (db.User, error) {
	users, _ := db.UserQuery(ctx)
	if len(users) == 0 {
		return db.User{}, errors.New(fmt.Sprintf("find no user, userCode: %s", code))
	}
	user := users[0]
	if user.Password != password {
		return db.User{}, errors.New(fmt.Sprintf("password error, userCode: %s", code))
	}
	return user, nil
}
