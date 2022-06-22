package admin

// 用户管理

import (
	"github.com/gin-gonic/gin"
	"goapi/comm/errno"
	"goapi/db/dao"
	"net/http"
)

type userRsp struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码md5
}

// 获取所有用户
func allUserHandler(c *gin.Context) {

	dbValue, err := dao.GetAllUser()
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}

	// 定义切片（即动态数组）
	var list []userRsp
	for _, v := range dbValue {
		list = append(list, userRsp{
			Username: v.Username,
			Password: v.Password,
		})
	}

	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{
		"list": list,
	}))
}
