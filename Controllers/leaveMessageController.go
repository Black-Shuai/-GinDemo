package Controllers

import (
	"GinDemo/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查找全部文章
func FindAllLeaveMessage(ctx *gin.Context)  {
	result,err :=Services.FindLeaveMessageService()
	fmt.Println(result)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取错误",
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}
}
