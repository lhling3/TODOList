package controller

import (
	"TODOList/service"
	"TODOList/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var UserClient = &User{service: &service.UserService{}}

type User struct {
	service *service.UserService
}

// @Summary      create user
// @Description  create user
// @Accept       json
// @Produce      json
// @Param        id body uint64 true "require id"
// @Param        name body string true "require name"
// @Param        password body string true "require password"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /admin/createuser [post]
func (u User) Create(ctx *gin.Context) {
	user := &store.User{}
	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数校验失败"})
		return
	}
	id, err := u.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "新增用户成功",
		"data": id,
	})
}

// @Summary      delete user
// @Description  delete user
// @Accept       json
// @Produce      json
// @Param        user_id query uint64 true "require user_id"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /admin/deleteuser [delete]
func (u User) Delete(ctx *gin.Context) {
	var idStr = ctx.Query("user_id")
	userId, err := strconv.ParseUint(idStr, 10, 64)
	//err不为空表示传入的参数不是uint，userId == 0表示没有指定删除用户Id，userId == 1表示管理员，不能删除
	if err != nil || userId == 0 || userId == 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "删除的用户ID不合法"})
		return
	}
	err = u.service.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除用户成功",
	})
}

// @Summary      update user password
// @Description  update user password
// @Accept       json
// @Produce      json
// @Param        id body uint64 true "require user_id"
// @Param        password body string true "require password"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /admin/updateuser [put]
func (u User) UpdatePassword(ctx *gin.Context) {
	user := &store.User{}
	err := ctx.ShouldBind(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数不合法"})
		return
	}
	if user.Id == 0 || user.Id == 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "修改的用户ID不合法"})
		return
	}
	if user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码参数不合法"})
		return
	}
	err = u.service.UpdatePassword(user.Id, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "修改用户密码成功",
	})
}
