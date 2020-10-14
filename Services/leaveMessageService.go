package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
	"fmt"
)

//查找留言内容
func FindLeaveMessageService()(lea []Models.LeaveMessage,err error)  {
	lea,err=Dao.FindLeaveMessageMapper()
	fmt.Println(lea)
	return
}
