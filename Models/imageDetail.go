package Models

type ImageDetail struct {
	Id int `json:"id"`
	ListId int `json:"list_id"`
	Detail string `json:"detail"`
	Url string `json:"url"`
	Author string `json:"author"`
	Createtime string `json:"createtime"`
	Updatetime string `json:"updatetime"`
}

func (ImageDetail)TableName() string  {
	return "tb_imagedetail"

}