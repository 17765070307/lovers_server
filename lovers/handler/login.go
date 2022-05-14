package handler

import (
	"Lovers/db"
	"Lovers/server/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginResp struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data db.User `json:"data"`
}

func LoginHandler(ctx *gin.Context) {
	code := ctx.PostForm("user_code")
	password := ctx.PostForm("user_password")
	user, _ := user.Login(ctx, code, password)
	ctx.JSON(http.StatusOK, LoginResp{Code: 0, Msg: "success.", Data: user})
}
