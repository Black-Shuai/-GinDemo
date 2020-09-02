package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
)

//向文章添加内容
func InsertArticleMapper(article Models.Article)(err error)  {
	//Mysql.DB.Raw("insert into tb_article (title,content,created_time) values('test','testcontent','2020-08-28 11:33:28')")
	dberr:=Mysql.DB.Create(&article)
	err=dberr.Error
	return
}
//查找全部文章
func FindArticleMapper()(article []Models.Article,err error)  {
	Mysql.DB.Find(&article)
	return
}
//查找文章分类
func FindArticleSortMapper(generalsort string)(article []Models.ArticleSort,err error)  {
	Mysql.DB.Find(&article).Where("general_sort=?",generalsort)
	return
}
//查找文章的大体分类
func FindGeneralSortMapper()(general []Models.GeneralSort,err error)  {
	Mysql.DB.Find(&general)
	return
}
