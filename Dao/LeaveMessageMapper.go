package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
)

//查找留言内容
func FindLeaveMessageMapper()(lea []Models.LeaveMessage,err error)  {
	Mysql.DB.Order("create_time desc").Find(&lea)
	for i:=0;i<len(lea);i++{
		Mysql.DB.Preload("LeaveAnswer").Find(&lea[i])
	}
	return
}
