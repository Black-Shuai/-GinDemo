package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
)

func FindImageListService() (imagelist[] Models.ImageList,err error)  {
	imagelist,err=Dao.FindImageListMapper()
	return
}
