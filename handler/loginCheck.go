package handler

import (
	"github.com/gin-gonic/gin"
	"db/redis"
	"utils"
)

func LoginMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.PostForm("userId")
		userToken := ctx.PostForm("token")

		res := new(utils.Response)
		if userId == "" || userToken == "" {
			res.FailCode(402, "用户未登录")
			ctx.JSON(200, res)
			ctx.Abort()
			return
		}

		token, err := redis.Get(userId)
		if err != nil {
			res.FailCode(402, "用户未登录")
			ctx.JSON(200, res)
			ctx.Abort()
			return
		}
		if token == "" || token != userToken {
			res.FailCode(402, "用户未登录")
			ctx.JSON(200, res)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
