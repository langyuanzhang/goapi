package dao

import (
	"goapi/comm/encrypt"
	"goapi/comm/log"
	"goapi/db"
	"goapi/db/model"
	"gorm.io/gorm"
	"time"
)

const userTableName = "user"

// AddUserRecordIfNeeded 增加用户，若username重复，则不做任何事
func AddUserRecordIfNeeded(username string, password string) error {
	cli := db.Get()
	var record *model.UserRecord
	if result := cli.Table(userTableName).
		Where("username = ?", username).
		First(&record); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			md5Pwd := encrypt.GenerateMd5(password)
			nowTime := time.Now()
			newUser := model.UserRecord{Username: username, Password: md5Pwd, CreateTime: nowTime, UpdateTime: nowTime}
			log.Debug(newUser)
			if err := cli.Table(userTableName).
				Create(&newUser).Error; err != nil {
				return err
			}
			log.Infof("Save User: %v", record)
		}
		return result.Error
	} else {
		log.Infof("User Already Exists: %v", record)
	}
	return nil
}

// GetAllUser 获取所有用户
func GetAllUser() ([]*model.UserRecord, error) {
	var records []*model.UserRecord
	cli := db.Get()
	result := cli.Table(userTableName)
	result = result.Find(&records)
	return records, result.Error
}
