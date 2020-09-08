package Services

import (
	"GinDemo/Dao"
	"GinDemo/Models"
	"regexp"
)

func InsertArticle(article Models.Article)(art Models.Article,err error)  {
	art,err=Dao.InsertArticleMapper(article)
	return art,err
}
func InsertContent(contentx []Models.TbContent) error  {
	return Dao.InsertContentMapper(contentx)
}
//查找全部文章
func FindAllArticle() (article []Models.Article,err error)  {
	return Dao.FindArticleMapper()
}
//查找全部文章
func FindArticleByIdMapper(articleid string) (article []Models.Article,err error)  {
	return Dao.FindArticleByIdMapper(articleid)
}
//查找文章具体分类
func FindAllArticleSort(generalsort string) (articlesort []Models.ArticleSort,err error)  {
	return Dao.FindArticleSortMapper(generalsort)
}
//查找文章大体分类
func FindAllGeneral() (general []Models.GeneralSort,err error)  {
	return Dao.FindGeneralSortMapper()
}
func Interception(cont string,art int)(tbc []Models.TbContent)  {
	var tbcon Models.TbContent
	var start,end int
	end=15374*2
	var artlen=len([]rune(cont))
	var i=1
	for start=15374;start<artlen;start=start+15374{
		if end<len([]rune(cont)){
			tbcon.Number=i
			tbcon.ArticleId=art
			tbcon.Content=string([]rune(cont)[start:end])
			re, _ := regexp.Compile("'")
			tbcon.Content= re.ReplaceAllString(tbcon.Content, "''")
			tbc=append(tbc, tbcon)
		} else {
			tbcon.Number=i
			tbcon.ArticleId=art
			tbcon.Content=string([]rune(cont)[start:])
			re, _ := regexp.Compile("'")
			tbcon.Content= re.ReplaceAllString(tbcon.Content, "''")
			tbc=append(tbc, tbcon)
		}
		end=end+15374
		i++
	}
	return
}
