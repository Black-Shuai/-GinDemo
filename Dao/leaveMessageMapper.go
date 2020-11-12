package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
	"fmt"
	"strconv"
)

//查找留言内容
func FindLeaveMessageMapper()(lea []Models.LeaveMessage,err error)  {
	Mysql.DB.Order("create_time desc").Find(&lea)
	for i:=0;i<len(lea);i++{
		Mysql.DB.Preload("LeaveAnswer").Find(&lea[i])
	}
	return
}
//添加留言内容
func AddLeaveMessageMapper(message Models.LeaveMessage) bool {
	DBerr:=Mysql.DB.Create(&message)
	fmt.Println(DBerr.Error)
	if DBerr.Error != nil {
		return false
	}else {
		return true
	}
}
//添加留言回答内容
func AddLeaveAnswerMapper(answer Models.LeaveAnswer) bool {
	DBerr:=Mysql.DB.Create(&answer)
	fmt.Println(answer.Id)
	var message Models.LeaveMessage
	Mysql.DB.Where("id=?",answer.Next_id).First(&message)
	fmt.Println(message)
	derr:=Mysql.DB.Model(&message).Update("next_id",strconv.Itoa(answer.Id))
	if DBerr.Error != nil&&derr.Error!=nil {
		return false
	}else {
		return true
	}
}
