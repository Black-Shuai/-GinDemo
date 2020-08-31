package Models

type ArticleSort struct {
	Id int `json:"id"`
	ArticleSort string `json:"article_sort"`
	GeneralSort string `json:"general_sort"`
}

func (ArticleSort)TableName() string {
	return "db_articlesort"
}