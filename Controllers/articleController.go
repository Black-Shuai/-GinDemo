package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(ctx *gin.Context)  {
	var article Models.Article
	err:=ctx.ShouldBindJSON(&article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dberr := Services.InsertArticle(article)
	if dberr!=nil {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":0,
			"Message":"数据插入错误",
			"Result":dberr,
		})
	}else {
		ctx.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据插入成功",
			"Result":"Success",
		})
	}
}
