package Models

type ImageList struct {
	Id int `json:"id"`
	ImageName string `json:"image_name"`
	UpdateTime string `json:"update_time"`
}
func (ImageList)TableName() string  {
	return "tb_imagelist"
}