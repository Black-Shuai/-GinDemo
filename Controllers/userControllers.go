package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Services"
	"GinDemo/Services/ServiceModels"
	"GinDemo/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
	//"fmt"
)
//声明相对应的结构体
var userlogin ServiceModels.User
var resResponse Util.ResultResponse
var user Models.User
//var userList []Models.User
//用户查找
func FindUser(c *gin.Context) {
	//数据解析，将json格式数据解析为结构体
	err := c.ShouldBindJSON(&userlogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user,err=Services.FindUser(userlogin)
	if err != nil||user.Id==""{
		c.JSON(http.StatusOK,resResponse.ErrResult())
		return
	}else {
		c.JSON(http.StatusOK,resResponse.SuccessResult(user))
	}
}
//查找全部用户
func FindAllUser(c *gin.Context) {
	userList,err:=Services.FindAllUser()
	if err != nil{
		c.JSON(http.StatusOK,resResponse.ErrResult())
		return
	}else {
		c.JSON(http.StatusOK,resResponse.SuccessResultList(userList))
	}
}
//查找登录用户
func Login(c *gin.Context) {
	Util.InitTime()
	err :=c.ShouldBindJSON(&user)
	fmt.Println(user)
	userList,err:=Services.Login(user)
	if err != nil||userList.Id==""{
		c.JSON(http.StatusOK,resResponse.ErrResult())
		return
	}else {
		c.JSON(http.StatusOK,resResponse.SuccessResultList1(userList))
	}
}