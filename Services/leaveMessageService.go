package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
	"fmt"
)

//查找留言内容
func FindLeaveMessageService()(lea []Models.LeaveMessage,err error)  {
	lea,err=Dao.FindLeaveMessageMapper()
	return
}
//添加留言
func AddLeaveMessageService(message Models.LeaveMessage) bool {
	fmt.Println(message)
	return Dao.AddLeaveMessageMapper(message)
}
