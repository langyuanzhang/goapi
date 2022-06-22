package admin

import (
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(e *gin.RouterGroup) {
	e.GET("/ping", pingHandler)
	e.Any("/test", testHandler)
	e.GET("/all_user", allUserHandler)

	// // 分组路由
	//g := e.Group("/admin", middleware.JWTMiddleWare)
	//g.POST("/username", testHandler)

}
