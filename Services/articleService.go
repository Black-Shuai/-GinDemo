package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
)

func InsertArticle(article Models.Article)(err error)  {
	return Dao.InsertArticleMapper(article)
}
//查找全部文章
func FindAllArticle() (article []Models.Article,err error)  {
	return Dao.FindArticleMapper()
}
//查找文章具体分类
func FindAllArticleSort(generalsort string) (articlesort []Models.ArticleSort,err error)  {
	return Dao.FindArticleSortMapper(generalsort)
}
//查找文章大体分类
func FindAllGeneral() (general []Models.GeneralSort,err error)  {
	return Dao.FindGeneralSortMapper()
}

