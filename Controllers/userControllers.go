package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Util"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
	//"fmt"
	"GinDemo/Services"
)
//用户查找
func UserFind(c *gin.Context) {
	var userService Services.User
	var resResponse Util.ResultResponse
	var user Models.User
	//数据解析，将json格式数据解析为结构体
	err := c.ShouldBindJSON(&userService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user,err=userService.Find()
	if err != nil||user.Id==""{
		c.JSON(http.StatusOK,resResponse.ErrResult())
		return
	}else {
		c.JSON(http.StatusOK,resResponse.SuccessResult(user))
	}
}