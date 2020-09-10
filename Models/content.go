package Models

type TbContent struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Number int `json:"number"`
	ArticleId int `json:"article_id"`
}
func (TbContent)TableName() string {
	return "tb_content"
}
