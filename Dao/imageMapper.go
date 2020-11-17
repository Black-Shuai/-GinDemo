package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
)
//查找图片分类
func FindImageListMapper()(imagelist[] Models.ImageList,err error)  {
	Mysql.DB.Find(&imagelist)
	return
}

//查找图片分类及信息
func FindImageMapper()(imagelist[] Models.ImageList,err error)  {
	Mysql.DB.Find(&imagelist)
	for i:=0;i<len(imagelist);i++ {
		imagelist[i].ImageDetail,err=FindImageDetailMapper(imagelist[i].Id)
	}
	return
}

//查找图片详情
func FindImageDetailMapper(list_id int)(imageDetail[] Models.ImageDetail,err error)  {
	Mysql.DB.Where("list_id=?",list_id).Find(&imageDetail)
	return
}
//添加图片信息
func AddImageDetailMapper(imagedetail Models.ImageDetail)(err error)  {
	Mysql.DB.Create(&imagedetail)
	return
}
//添加图片列表
func AddImageListMapper(list Models.ImageList)(err error)  {
	Mysql.DB.Create(&list)
	return
}

