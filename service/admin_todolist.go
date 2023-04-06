package service

import (
	"TODOList/store"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminTodoListService struct{}

func (s AdminTodoListService) Update(ctx *gin.Context, todoList *store.TodoList) error {
	//根据用户id得到当前用户的待办事项map
	userTodoList, ok := store.TODOListMap[todoList.UserID]
	if !ok {
		return errors.New("该用户尚未创建任何待办事项，请先创建")
	}
	_, ok = userTodoList[todoList.Id]
	if !ok {
		return errors.New("该用户不存在该代办事项，无法更改")
	}
	//将新的代办事项存入当前用户的map中
	todoList.UpdateTime = time.Now()
	userTodoList[todoList.Id] = *todoList
	return nil
}

func (s AdminTodoListService) Delete(ctx *gin.Context, todoList *store.TodoList) error {
	userTodoList, ok := store.TODOListMap[todoList.UserID]
	if !ok {
		return errors.New("当前用户没有可删除的待办事项")
	}
	delete(userTodoList, todoList.Id)
	return nil
}

func (s AdminTodoListService) QueryByAdmin(ctx *gin.Context) (map[uint64]map[uint]store.TodoList, error) {
	return store.TODOListMap, nil
}
