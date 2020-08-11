package main

import (
	"fmt"
	"GinDemo/entity"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context)  {
	fmt.Println("Hello Golang")
	//
	c.JSON(200,gin.H{
		"message":"Hello Golang",
	})
}
func main()  {

	//返回默认路由引擎
	r := gin.Default()
	//通过GET方法调用test函数
	r.GET("/hello",test)
	//RESTful API设计
	//获取数据, 匹配的url格式:  /select?firstname=Jane&lastname=Doe
	r.GET("/select", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
		res := entity.SuccessResponse{
			Code: 0,
			Msg: "获取数据",
			Data:entity.ResultData{
				Id:0,
				Firstname: firstname,
				Lastname: lastname,
			},
		}
		context.JSON(200,res)
	})
	//添加数据接口
	r.POST("/insert", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	//更新数据接口
	r.PUT("/update/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(200,gin.H{
			"message":"更新数据成功",
			"name":name,
		})
	})
	//删除数据接口
	r.DELETE("/delete", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"删除数据成功",
		})
	})
	//启动
	r.Run()
}
