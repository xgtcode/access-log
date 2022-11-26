package service

import (
	"access-log-app/pkg/global"
	"access-log-app/pkg/model"
	log "github.com/xgtcode/log-demo"
)

func Visit(user string) (err error){
	// 将记录添加到数据库中  如果没有该user,就直接插入
	// 如果有就直接+1
	var accessLogs []model.AccessLog
	err = global.GlobalDb.Model(model.AccessLog{}).Where("user = ?", user).Find(&accessLogs).Error
	if err != nil{
		return err
	}
	if len(accessLogs) == 0{
		err = global.GlobalDb.Create(&model.AccessLog{
			User: user,
			Count: 1,
		}).Error
		return err
	}
	err = global.GlobalDb.Model(accessLogs).Where("user = ?", user).Update("count", accessLogs[0].Count+1).Error
	return err
}

func CountVisitLog(user string) (int, error){
	accessLog := &model.AccessLog{}
	err := global.GlobalDb.FirstOrInit(accessLog, &model.AccessLog{User: user, Count: 0}).Error
	if err != nil{
		return 0, err
	}
	log.Infof("user is %s, count is %d", user, accessLog.Count)
	return accessLog.Count, nil
}
