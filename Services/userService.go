package Services

import (
	"GinDemo/Models"
)

type User struct {
	UserName string
	UserLoginName	string
	UserPassword string
}

func (this *User)Find()(user Models.User,err error)  {
	user.UserName=this.UserName
	user.UserPassword=this.UserPassword
	user,err =user.Find()
	return
}
