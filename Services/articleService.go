package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
)

func InsertArticle(article Models.Article)(err error)  {
	return Dao.InsertArticleMapper(article)
}
