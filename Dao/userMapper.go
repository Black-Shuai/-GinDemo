package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
	"GinDemo/Services/ServiceModels"
	_ "github.com/lib/pq"
)


func FindUser(user2 ServiceModels.User)(user Models.User,err error)  {
	Mysql.DB.Where("user_login_name=? and user_password=?",user2.UserLoginName,user2.UserPassword).First(&user)
	return
}

func FindAllUser()(user []Models.User,err error)  {
	Mysql.DB.Find(&user)
	return
}
func FindAlUser()(user []Models.User,err error)  {
	return
}
