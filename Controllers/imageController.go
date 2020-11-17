package Controllers

import (
	"GinDemo/Models"
	"GinDemo/Services"
	"GinDemo/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//查找图片分类数据
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

//查找图片大致信息
func FindImageController(c *gin.Context)  {
	imageList,err:=Services.FindImageService()
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

//查找图片详细信息
func FindImageDetailController(c *gin.Context)  {
	listid := c.Query("listid")
	result,err :=Services.FindImageDetailService(listid)
	if err==nil {
		c.JSON(http.StatusOK,gin.H{
			"Code":1,
			"Message":"数据获取成功",
			"Data":result,
		})
	}else {
		c.JSON(http.StatusBadRequest,gin.H{
			"Code":0,
			"Message":"数据获取失败",
		})
	}

}

//添加图片详细信息
func AddImageDetailController(c *gin.Context)  {
	var imagedetail[] Models.ImageDetail
	err:=c.ShouldBindJSON(&imagedetail)
	if err == nil {
		for i:=0;i< len(imagedetail);i++ {
			//获取当前时间，并赋值给结构体
			imagedetail[i].Createtime=Util.InitTime()
			imagedetail[i].Updatetime=Util.InitTime()
			Serr:=Services.AddImageDetailService(imagedetail[i])
			if Serr == nil {
				c.JSON(http.StatusOK,gin.H{
					"Code":1,
					"Message":"数据插入成功",
				})
			}else {
				c.JSON(http.StatusBadRequest,gin.H{
					"Code":0,
					"Message":"数据插入失败",
				})
			}
		}
	}else {
		fmt.Println(err)
	}

}

//添加图片分类信息
func AddImageListController(c *gin.Context)  {
	var imageList Models.ImageList
	err:=c.ShouldBindJSON(&imageList)
	if err == nil {
			imageList.UpdateTime=Util.InitTime()
			Serr:=Services.AddImageListService(imageList)
			if Serr == nil {
				c.JSON(http.StatusOK,gin.H{
					"Code":1,
					"Message":"数据插入成功",
				})
			}else {
				c.JSON(http.StatusBadRequest,gin.H{
					"Code":0,
					"Message":"数据插入失败",
				})
			}
	}
}


