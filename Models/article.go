package Models

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ArticleSort string `json:"article_sort"`
	Content string `json:"content"`
	Author string `json:"author"`
	Background string `json:"background"`
	CreatedTime string `json:"created_time"`
}

func (Article)TableName() string  {
	return "tb_article"
}
