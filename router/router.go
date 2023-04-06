package router

import (
	"TODOList/controller"
	_ "TODOList/docs"
	"TODOList/middleware"
	"TODOList/store"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	sessionMap := store.NewSessionMap()
	r := gin.Default()
	r.POST("/login", middleware.Login(sessionMap))
	r.GET("/logout", middleware.Logout(sessionMap))

	//一般用户组
	v1 := r.Group("/user")
	v1.Use(middleware.UserAuth(sessionMap))
	{
		v1.POST("/create", controller.TodoListClient.Create)
		v1.PUT("/update", controller.TodoListClient.Update)
		v1.DELETE("/delete", controller.TodoListClient.Delete)
		v1.GET("/querybyuser", controller.TodoListClient.QueryByUser)
	}

	//管理员组
	v2 := r.Group("/admin")
	v2.Use(middleware.UserAuth(sessionMap))
	v2.Use(middleware.AdminAuth(sessionMap))
	{
		v2.POST("/createuser", controller.UserClient.Create)
		v2.PUT("/updateuser", controller.UserClient.UpdatePassword)
		v2.DELETE("/deleteuser", controller.UserClient.Delete)

		v2.PUT("/updatelist", controller.AdminTodoListClient.Update)
		v2.DELETE("/deletelist", controller.AdminTodoListClient.Delete)
		v2.GET("/querybyadmin", controller.AdminTodoListClient.QueryByAdmin)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}
