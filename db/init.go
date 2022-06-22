package db

import (
	"fmt"
	"goapi/comm/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/patrickmn/go-cache"
	"goapi/comm/log"
)

var dbInstance *gorm.DB
var cacheInstance *cache.Cache

// Init 初始化数据库
func Init() error {
	var user, pwd, addr, dataBase string
	user = config.DbConf.User
	pwd = config.DbConf.Pwd
	addr = config.DbConf.Addr
	dataBase = config.DbConf.Database
	if dataBase == "" {
		dataBase = "goapi"
	}
	tcp := "tcp"
	source := "%s:%s@" + tcp + "(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	log.Debug("start inits mysql with ::::: " + source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})

	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish inits mysql with ", source)

	checkTables()

	// 初始化cache
	cacheInstance = cache.New(5*time.Minute, 10*time.Minute)

	return nil
}

func checkTables() {
	dbInstance.Exec("CREATE TABLE IF NOT EXISTS `user` ( `id` INT NOT NULL AUTO_INCREMENT, `username` VARCHAR(32) NOT NULL, `password` VARCHAR(64) NOT NULL, `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `updatetime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`ID`), UNIQUE KEY `user_username_uindex` (`username`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
}

// Get
func Get() *gorm.DB {
	return dbInstance
}

// GetCache
func GetCache() *cache.Cache {
	return cacheInstance
}
