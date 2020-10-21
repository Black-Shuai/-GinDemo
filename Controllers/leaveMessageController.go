package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Services"
	"GinDemo/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查找全部文章
func FindAllLeaveMessage(ctx *gin.Context)  {
	result,err :=Services.FindLeaveMessageService()
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
//添加留言信息
func AddLeaveMessage(ctx *gin.Context)  {
	var leaveMessage Models.LeaveMessage
	err:=ctx.ShouldBindJSON(&leaveMessage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"提交数据错误",
		})
	}else {
		//获取当前时间，并赋值给结构体
		leaveMessage.CreateTime=Util.InitTime()
		leaveMessage.NextId=""
		res:=Services.AddLeaveMessageService(leaveMessage)
		if res {
			ctx.JSON(http.StatusOK,gin.H{
				"Code":1,
				"Message":"留言添加成功",
			})
		}
	}

}
