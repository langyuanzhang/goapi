package config

import (
	"github.com/go-ini/ini"
	"goapi/comm/encrypt"
	"goapi/comm/log"
)

// Server 系统配置结构体
type Server struct {
	JwtSecret     string
	JwtIssue      string
	JwtExpireTime int32
	AesKey        string
}

// DBconf 数据库配置结构体
type DBconf struct {
	User     string
	Pwd      string
	Addr     string
	Database string
}

// Comm 常规配置结构体
type Comm struct {
	Version string
}

// Logconf 日志配置结构体
type Logconf struct {
	Debug bool
	Path  string
}

var ServerConf = &Server{}
var CommConf = &Comm{}
var DbConf = &DBconf{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("comm/config/server.conf")
	if err != nil {
		log.Errorf("load server.conf': %v", err)
		return
	}
	mapTo("server", ServerConf)
	mapTo("comm", CommConf)
	mapTo("db", DbConf)
	if ServerConf.AesKey == "" {
		ServerConf.AesKey = encrypt.GenerateMd5("zhangzhang@2B2")
	}
	if ServerConf.JwtSecret == "" {
		ServerConf.JwtSecret = encrypt.GenerateMd5("zhangzhang@2B2")
	}
	log.Info(ServerConf)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Errorf("%s err: %v", section, err)
	}
}
