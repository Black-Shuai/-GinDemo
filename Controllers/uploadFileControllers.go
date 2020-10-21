package Controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Upload(ctx *gin.Context) {
	//获取文件头
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	var fName=time.Now().Format("2006-01")//获取当前时间以便于文件创建
	_,ferr := os.Stat("file/"+fName)//检查文件夹是否存在
	if ferr == nil {
	}
	if ferr!=nil {//文件夹不存在则创建新的文件夹
		os.MkdirAll("file/"+fName, os.ModePerm)
	}
	//保存文件到服务器本地
	//存储与项目文件夹下的file文件夹下
	filePath :="file/"+fName+"/"+fileName
	//SaveUploadedFile(文件头，保存路径)
	xerr := ctx.SaveUploadedFile(file, filePath);
	if  xerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code":0,
			"Message":"上传文件失败",
			"error":xerr.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Code":1,
		"Message":"上传文件成功",
		"urlPath":"http://localhost:8088/api/file/getimage?imageName="+fName+"/"+fileName,
	})
}
func GetImage(c *gin.Context){
	imageName := c.Query("imageName")
	path:="file/"+imageName
	file, _ := ioutil.ReadFile(path)
	c.Writer.WriteString(string(file))
}