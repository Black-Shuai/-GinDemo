package Models

type TestUser struct {
	Id 				int64 `json:"id"`
	OpenId 			string `json:"open_id"`
	UserName 		string `json:"user_name"`
	UserSex 		string `json:"user_sex"`
	UserMobile 		string `json:"user_mobile"`
	UserLoginName	string `json:"user_login_name"`
	UserPassword 	string `json:"user_password"`
	CreateTime 		string `json:"create_time"`
	UpdateTime 		string `json:"update_time"`
	IsTrue			string `json:"is_true"`
}
