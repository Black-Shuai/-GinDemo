package Models

import (
	Mysql "GinDemo/Databases"
)

type User struct {
	Id 				string `json:"id"`
	OpenId 			string `json:"open_id"`
	UserName 		string `json:"user_name"`
	UserSex 		string `json:"user_sex"`
	UserMobile 		string `json:"user_mobile"`
	UserLoginName	string `json:"user_login_name"`
	UserPassword 	string `json:"user_password"`
	CreateTime 		string `json:"create_time"`
	UpdateTime 		string `json:"update_time"`
}

//设置Test的表名为`tb_user`
func (User) TableName() string {
	return "tb_user"
}

func (this *User)Find()(user User,err error)  {
	Mysql.DB.Where("user_name=? and user_password=?",this.UserName,this.UserPassword).First(&user)
	return
}
