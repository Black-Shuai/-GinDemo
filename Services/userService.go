package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
	"GinDemo/Services/ServiceModels"
)

func FindUser(user2 ServiceModels.User)(user Models.User,err error)  {
	//对controller接受到的数据进行匹配
	user,err =Dao.FindUser(user2)
	return
}
func FindAllUser()(user []Models.User,err error)  {
	//对controller接受到的数据进行匹配
	user,err =Dao.FindAllUser()
	return
}
