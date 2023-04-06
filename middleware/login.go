package middleware

import (
	"TODOList/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

// @Summary      user login
// @Description  user login
// @Accept       json
// @Produce      json
// @Param        id query uint64 true "require id"
// @Param        name query string true "require name"
// @Param        password query string true "require password"
// @Success      200 {object} 成功登陆
// @Failure      400 {object} 用户ID不合法,用户不存在,用户名或密码错误
// @Router       /login [post]
func Login(sessionMap *store.SessionMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		user, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户ID不合法"})
		}
		userName := c.Query("name")
		password := c.Query("password")
		// 校验是否合法用户
		v, ok := store.UserMap[user]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			return
		}
		if userName != v.Name || password != v.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
			return
		}

		// 合法用户，成功登陆，往客户端写cookie并创建一条对应session
		sessionId := uuid.New().String()
		c.SetCookie(sessionKey, sessionId, int(time.Hour.Seconds()), "/", "localhost", false, true)
		sessionMap.Put(sessionId, v)
		c.JSON(http.StatusOK, gin.H{"message": "成功登陆"})
	}
}
