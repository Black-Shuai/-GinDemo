package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
	"GinDemo/Services/ServiceModels"
	_ "github.com/lib/pq"
)

type Resultq struct {
	Id string `json:"id"`
	UserName string `json:"user_name"`
	Sex string `json:"sex"`
}

func FindUser(user2 ServiceModels.User)(user Models.User,err error)  {
	Mysql.DB.Where("user_login_name=? and user_password=?",user2.UserLoginName,user2.UserPassword).Preload("Sex").First(&user)
	return
}

func FindAllUser()(user []Models.User,err error)  {
	Mysql.DB.Find(&user)
	return
}

func Login(user3 Models.User)  (res Resultq,err error) {
	Mysql.DB.Raw("select a.id,a.user_name,b.name as sex from tb_user as a,db_sex as b where a.user_sex=b.id and a.user_login_name=? and a.user_password=?",user3.UserLoginName,user3.UserPassword).Scan(&res)
	return
}
