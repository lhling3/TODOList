package middleware

import (
	"TODOList/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户登出
// @Summary      user logout
// @Description  user logout
// @Produce      json
// @Success      200 {object} string
// @Failure      400 {object} string
// @Router       /logout [get]
func Logout(sessionMap *store.SessionMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie(sessionKey)
		if err != nil {
			// http request中不存在这个cookie，用户未登陆
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "未登陆用户不能登出"})
			return
		}

		sessionMap.Del(sessionId)
		// 通过把MaxAge设置为-1来删除客户端cookie
		c.SetCookie(sessionKey, sessionId, -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "成功登出"})
	}
}
