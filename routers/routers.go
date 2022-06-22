package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi/api/admin"
	"goapi/middleware"
)

type Option func(*gin.RouterGroup)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogMiddleWare)

	// 路由
	Include(admin.Routers)
	g := r.Group("/api")
	for _, opt := range options {
		opt(g)
	}

	// 静态文件
	g.Static("/assets", "client/assets")
	r.LoadHTMLGlob("client/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "404 not found",
		})
	})
	return r
}
