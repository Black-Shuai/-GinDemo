package Models

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ArticleSort string `json:"article_sort"`
	Content string `json:"content"`
	CreatedTime string `json:"created_time"`
}

func (Article)TableName() string  {
	return "tb_article"
}
