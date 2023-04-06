package controller

import (
	"TODOList/service"
	"TODOList/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AdminTodoListClient = &AdminTodoList{service: &service.TodoListService{}}

type AdminTodoList struct {
	service *service.TodoListService
}

// @Summary      update user todoList
// @Description  update user todoList
// @Accept       json
// @Produce      json
// @Param        id body uint true "require id"
// @Param        title body string true "require tiele"
// @Param        body body string true "require body"
// @Param        user_id body uint64 true "require user_id"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /admin/updatelist [put]
func (t AdminTodoList) Update(ctx *gin.Context) {
	todoList := &store.TodoList{}
	if err := ctx.ShouldBind(todoList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数校验失败"})
		return
	}
	err := t.service.Update(ctx, todoList)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "更新待办事项失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功更新一条代办事项",
	})
}

// @Summary      delete user todoList
// @Description  delete user todoList
// @Accept       json
// @Produce      json
// @Param        id query uint true "require id"
// @Param        user_id query uint64 true "require user_id"
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /admin/deletelist [delete]
func (t AdminTodoList) Delete(ctx *gin.Context) {
	todoList := &store.TodoList{}
	if err := ctx.ShouldBind(todoList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数校验失败,请指定要删除的待办事项"})
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

// @Summary      query all user todoList
// @Description  query all user todoList
// @Produce      json
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /admin/querybyadmin [get]
func (t AdminTodoList) QueryByAdmin(ctx *gin.Context) {
	//查询所有用户的所有待办事项
	list, err := t.service.QueryByAdmin(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "查询所有待办事项失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询所有用户待办事项成功",
		"data": list,
	})
}
