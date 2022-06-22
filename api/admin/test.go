package admin

// 人员管理

import (
	"github.com/gin-gonic/gin"
	"goapi/comm/errno"
	"goapi/comm/log"
	"net/http"
)

// PING
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"errorMsg": "pong",
	})
}

// 测试
func testHandler(c *gin.Context) {

	name, exists := c.GetQuery("name")
	log.Info("exists is ", exists)
	if !exists {
		log.Error("name not exists ")
		c.JSON(http.StatusOK, errno.ErrInvalidParam)
		return
	}
	log.Info("name is ", name)
	c.JSON(http.StatusOK, errno.OK)
}
