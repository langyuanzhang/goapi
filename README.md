# goapi
[![GitHub license](https://img.shields.io/github/license/WeixinCloud/wxcloudrun-wxcomponent)](https://goapi)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/WeixinCloud/wxcloudrun-wxcomponent)

学习版go语言API接口，改编自微信第三方平台管理工具模版

## 功能介绍
此项目提供最简易的go语言web接口入门学习，可参考完整版[微管家](github.com/WeixinCloud/wxcloudrun-wxcomponent)

#### 主要第三方组件
- web框架: github.com/gin-gonic/gin
- 日志: github.com/rs/zerolog
- 数据库驱动: gorm.io/driver/mysql
- ORM: gorm.io/gorm
- 缓存: github.com/patrickmn/go-cache
- JWT: github.com/golang-jwt/jwt/v4

## 目录结构
```
.
├── Dockerfile
├── README.md
├── api                                 // 后端api
│   └── admin                           // 后台管理路由控制器部分
├── client                              // 前端
│   ├── assets                          // 静态资源
│   └── index.html
├── comm                                // 后端公共模块
│   ├── config                          // 配置
│   ├── encrypt                         // 加密
│   ├── errno                           // 错误码
│   ├── httputils                       // http
│   ├── inits                           // 初始化
│   ├── log                             // 日志
│   └── utils                           // 其他工具
├── db                                  // 数据库相关
│   ├── dao
│   ├── init.go
│   └── model
├── go.mod
├── go.sum
├── main.go
├── middleware                          // 中间件
│   ├── jwt.go                          // jwt
│   └── log.go                          // 日志
└── routers                             // 路由
    └── routers.go

```

## 其他说明
#### 本地调试
服务启动前会从环境变量中读取数据库配置，自行写入环境变量后运行一下代码，即可在本地启动服务。
```
go run main
```

#### 数据表
```
+-----------------------+
| Tables_in_goapi |
+-----------------------+
| user                  |
+-----------------------+
```
- user: 用户表


## 部署配置
需复制「example.server.conf」重命名「server.conf」，并按实际情况配置才可正常使用

## License

[MIT](./LICENSE)
