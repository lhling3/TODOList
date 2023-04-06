package controller

import (
	"TODOList/service"
	"TODOList/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

var TodoListClient = &TodoList{service: &service.TodoListService{}}

type TodoList struct {
	service *service.TodoListService
}

// @Summary      create user todoList
// @Description  create user todoList
// @Accept       json
// @Produce      json
// @Param        id body uint true "require id"
// @Param        title body string true "require tiele"
// @Param        body body string true "require body"
// @Param        user_id body uint64 false " user_id"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /user/create [post]
func (t TodoList) Create(ctx *gin.Context) {
	todoList := &store.TodoList{}
	if err := ctx.ShouldBind(todoList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数校验失败"})
		return
	}
	//从上下文中取出当前用户信息
	user, _ := ctx.Get("user")
	userData := user.(*store.User)
	todoList.UserID = userData.Id

	id, err := t.service.Create(ctx, todoList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "新增待办事项失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "成功创建一条代办事项",
		"data": id,
	})
}

// @Summary      update user todoList
// @Description  update user todoList
// @Accept       json
// @Produce      json
// @Param        id body uint true "require id"
// @Param        title body string true "require tiele"
// @Param        body body string true "require body"
// @Param        user_id body uint64 false " user_id"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /user/update [put]
func (t TodoList) Update(ctx *gin.Context) {
	todoList := &store.TodoList{}
	if err := ctx.ShouldBind(todoList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数校验失败"})
		return
	}
	//从上下文中取出当前用户信息
	user, _ := ctx.Get("user")
	userData := user.(*store.User)
	todoList.UserID = userData.Id

	err := t.service.Update(ctx, todoList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "更新待办事项失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功更新一条代办事项",
	})
}

// @Summary      query user todoList
// @Description  query user todoList
// @Accept       json
// @Produce      json
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /user/querybyuser [get]
func (t TodoList) QueryByUser(ctx *gin.Context) {
	//从上下文中取出当前用户信息
	user, _ := ctx.Get("user")
	userData := user.(*store.User)
	//根据用户Id查询该用户的所有待办事项
	list, err := t.service.QueryByUser(ctx, userData.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "查询待办事项失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询当前用户待办事项成功",
		"data": list,
	})
}

// @Summary      delete user todoList
// @Description  delete user todoList
// @Accept       json
// @Produce      json
// @Param        id query uint true "require id"
// @Param        title query string false "require tiele"
// @Param        body query string false "require body"
// @Param        user_id query uint64 true " user_id"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /user/delete [delete]
func (t TodoList) Delete(ctx *gin.Context) {
	todoList := &store.TodoList{}
	if err := ctx.ShouldBind(todoList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数校验失败,请指定要删除的待办事项"})
		return
	}
	//从上下文中取出当前用户信息
	user, _ := ctx.Get("user")
	userData := user.(*store.User)
	if userData.Name != "admin" && userData.Id != todoList.UserID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "非管理员，只能删除自己的待办事项"})
		return
	}
	err := t.service.Delete(ctx, todoList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "删除待办事项失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功删除一条代办事项",
	})
}
