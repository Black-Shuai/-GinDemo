package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
)

func FindImageListMapper()(imagelist[] Models.ImageList,err error)  {
	Mysql.DB.Find(&imagelist)
	return
}
