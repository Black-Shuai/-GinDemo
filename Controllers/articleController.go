package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Services"
	"GinDemo/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(ctx *gin.Context)  {
	var article Models.Article
	err:=ctx.ShouldBindJSON(&article)
	//获取当前时间，并赋值给结构体
	article.CreatedTime=Util.InitTime()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dberr := Services.InsertArticle(article)
	if dberr!=nil {
		ctx.JSON(http.StatusAccepted,gin.H{
			"Code":0,
			"Message":"数据插入错误",
			"Data":dberr,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据插入成功",
			"Result":"Success",
		})
	}
}
//查找全部文章
func FindAllArticle(ctx *gin.Context)  {
	result,err :=Services.FindAllArticle()
	if err != nil {
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
