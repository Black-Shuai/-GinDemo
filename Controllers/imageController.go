package Controllers

import (
	"GinDemo/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)
func FindImageListController(c *gin.Context)  {
	imageList,err:=Services.FindImageListService()
	if err != nil {
		c.JSON(http.StatusOK,resResponse.ErrResult())
		return
	}else {
		c.JSON(http.StatusOK,gin.H{
			"Code": 1,
			"Message":"数据获取成功",
			"Data":imageList,
			})
	}
}
