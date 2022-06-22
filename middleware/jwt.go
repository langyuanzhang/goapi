package middleware

import (
	"net/http"
	"time"

	"goapi/comm/errno"
	"goapi/comm/log"

	"github.com/gin-gonic/gin"

	"goapi/comm/utils"
)

// JWTMiddleWare 中间件
func JWTMiddleWare(c *gin.Context) {
	code := errno.OK
	strToken := c.Request.Header.Get("Authorization")
	token := utils.GetToken(strToken)
	log.Debugf("jwt[%s]", token)

	var err error
	var claims *utils.Claims

	if token == "" {
		code = errno.ErrNotAuthorized
	} else {
		claims, err = utils.ParseToken(token)
		if err != nil {
			code = errno.ErrAuthTokenErr
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = errno.ErrAuthTimeout
		}
	}

	if code != errno.OK {
		c.JSON(http.StatusOK, code)
		c.Abort()
		return
	}

	log.Debugf("id:%s UserName:%s", claims.ID, claims.UserName)

	c.Set("jwt", claims)

	c.Next()
}
