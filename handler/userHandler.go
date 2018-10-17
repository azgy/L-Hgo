package handler

import (
	"model"
	"utils"
	"strconv"
	"github.com/gin-gonic/gin"
	"logger"
	"db/redis"
	"crypto/md5"
	"encoding/hex"
)

func Login(ctx *gin.Context) {
	user := new(model.User)
	res := new(utils.Response)
	if err := ctx.Bind(user); err != nil {
		res.Fail("传入的用户信息有误")
		ctx.JSON(200, res)
		return
	}
	md5Pass := md5.New()
	md5Pass.Write([]byte(user.Password))
	pass := hex.EncodeToString(md5Pass.Sum(nil))
	user.Password = pass
	has, err := user.Query()
	if err != nil {
		res.Fail(err.Error())
		ctx.JSON(200, res)
		return
	}
	if !has {
		res.Fail("用户名或密码错误")
		ctx.JSON(200, res)
		return
	}

	redis.Set(strconv.Itoa(user.Id), pass)

	ctx.SetCookie(strconv.Itoa(user.Id), pass, 0, "", "", false, true)

	res.Succ(user)
	ctx.JSON(200, res)
}

func AddUser(ctx *gin.Context) {
	user := new(model.User)
	res := new(utils.Response)
	if err := ctx.Bind(user); err != nil {
		res.Fail("传入的用户信息有误")
		ctx.JSON(200, res)
		return
	}

	md5Pass := md5.New()
	md5Pass.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(md5Pass.Sum(nil))
	_, err := user.Save()
	if err != nil {
		res.Fail(err.Error())
		ctx.JSON(200, res)
		return
	}

	res.Succ(user)
	ctx.JSON(200, res)
}

func GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	res := new(utils.Response)
	if err != nil {
		logger.Info("传入的id有误,id=%d", id)
		res.Fail("传入的id有误")
		ctx.JSON(200, res)
		return
	}
	user := &model.User{Id: id}
	has, err := user.Query()
	if err != nil {
		res.Fail(err.Error())
		ctx.JSON(200, res)
		return
	}
	if !has {
		res.Fail("没有该用户")
		ctx.JSON(200, res)
		return
	}

	res.Succ(user)
	ctx.JSON(200, res)
}
