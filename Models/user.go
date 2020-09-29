package Models

//与数据库表格字段相对应
type User struct {
	Id 				string `json:"id";gorm:"primary_key"`
	OpenId 			string `json:"open_id"`
	UserName 		string `json:"user_name"`
	UserMobile 		string `json:"user_mobile"`
	UserLoginName	string `json:"user_login_name"`
	UserPassword 	string `json:"user_password"`
	CreateTime 		string `json:"create_time"`
	UpdateTime 		string `json:"update_time"`
	UserSex 		string `json:"user_sex"`
	//根据所添加的外键，检索将对应的外键与相匹配的表格做连接
	Sex Sex `gorm:"foreignkey:UserSex;"`
}

//设置Test的表名为`tb_user`
func (User) TableName() string {
	return "tb_user"
}
