package middleware

import (
	"TODOList/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

var sessionKey = "session_id"

// 用户认证中间件
func UserAuth(sessionMap *store.SessionMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(sessionKey)
		if err != nil {
			// http request中不存在这个cookie，用户未登陆
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未登陆"})
			return
		}

		if !sessionMap.Exist(sessionId) {
			// cookie中的sessionId不是合法的
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "非法SessionID"})
			return
		}

		// 把对应用户的数据放在上下文中方便后续取用
		c.Set("user", sessionMap.Get(sessionId))

		c.Next()
	}
}

// 管理员认证中间件
func AdminAuth(sessionMap *store.SessionMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		userData := user.(*store.User)
		if userData.Name != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "非管理员，没有权限"})
			return
		}
		c.Next()
	}
}
