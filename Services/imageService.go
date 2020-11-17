package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
	"strconv"
)

func FindImageListService() (imagelist[] Models.ImageList,err error)  {
	imagelist,err=Dao.FindImageListMapper()
	return
}

func FindImageService() (imagelist[] Models.ImageList,err error)  {
	imagelist,err=Dao.FindImageMapper()
	return
}

func FindImageDetailService(list_id string) (imageDetail[] Models.ImageDetail,err error)  {
	listId,e:=strconv.Atoi(list_id)
	if e == nil {
		imageDetail,err=Dao.FindImageDetailMapper(listId)
	}
	return
}

func AddImageDetailService(imagedetail Models.ImageDetail) (err error)  {
	err=Dao.AddImageDetailMapper(imagedetail)
	return
}
func AddImageListService(list Models.ImageList) (err error)  {
	err=Dao.AddImageListMapper(list)
	return
}
