package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
	"fmt"
)

//查找留言内容
func FindLeaveMessageMapper()(lea []Models.LeaveMessage,err error)  {
	Mysql.DB.Order("create_time desc").Find(&lea)
	for i:=0;i<len(lea);i++{
		Mysql.DB.Preload("LeaveAnswer").Find(&lea[i])
	}
	return
}

func AddLeaveMessageMapper(message Models.LeaveMessage) bool {
	DBerr:=Mysql.DB.Create(&message)
	fmt.Println(DBerr.Error)
	if DBerr.Error != nil {
		return false
	}else {
		return true
	}
}
