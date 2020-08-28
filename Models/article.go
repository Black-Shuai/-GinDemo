package Models

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedTime string `json:"created_time"`
}

func (Article)TableName() string  {
	return "tb_article"
}
