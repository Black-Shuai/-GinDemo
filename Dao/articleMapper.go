package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
	"fmt"
)

//向文章添加内容
func InsertArticleMapper(article Models.Article)(art Models.Article,err error)  {
	//Mysql.DB.Raw("insert into tb_article (title,content,created_time) values('test','testcontent','2020-08-28 11:33:28')")
	dberr:=Mysql.DB.Create(&article)
	err=dberr.Error
	art=article
	//fmt.Println(art)
	return
}
//添加多余的文章
func InsertContentMapper(content []Models.TbContent) error {
	sql:="Insert into tb_content (`content`,`number`,`article_id`) values "
	for i,e:=range content{
		if len(content)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("('%s','%d','%d');",e.Content, e.Number, e.ArticleId,)
		}else {
			sql += fmt.Sprintf("('%s','%d','%d'),",e.Content, e.Number, e.ArticleId,)
		}
	}
	err:=Mysql.DB.Exec(sql)
	fmt.Println("true")
	return err.Error
}
//查找全部文章
func FindArticleMapper()(article []Models.Article,err error)  {
	Mysql.DB.Order("id desc").Find(&article)
	return
}//查找前6条文章
func FindArticle6()(article []Models.Article,err error)  {
	Mysql.DB.Raw("select a.id,b.article_sort as article_sort,a.title,a.content,a.background,a.created_time from tb_article as a,db_articlesort as b where a.article_sort=b.id order by a.id desc limit 6",).Scan(&article)
	return
}
//按照ID查找查找文章
func FindArticleByIdMapper(articleid string)(article []Models.Article,err error)  {
	var conx []Models.TbContent
	Mysql.DB.Where("id=?",articleid).Find(&article)
	total := 0
	Mysql.DB.Find(&conx).Where("article_id=?",articleid).Count(&total)
	if total!=0{
		for i:=1;i<=total;i++{
			var con Models.TbContent
			Mysql.DB.Where("number=? and article_id=?",i,articleid).Find(&con)
			article[0].Content=article[0].Content+con.Content
		}
	}
	return
}
//查找文章分类
func FindArticleSortMapper(generalsort string)(article []Models.ArticleSort,err error)  {
	Mysql.DB.Where("general_sort=?",generalsort).Find(&article)
	return
}
//查找文章的大体分类
func FindGeneralSortMapper()(general []Models.GeneralSort,err error)  {
	Mysql.DB.Find(&general)
	return
}

func FindArticlexMapper(id int)(article []Models.Article)  {
	var general Models.GeneralSort
	Mysql.DB.Where("id=?",id).Find(&general)
	var articlesort []Models.ArticleSort
	Mysql.DB.Where("general_sort=?",general.Id).Find(&articlesort)
	for i:=0; i<len(articlesort);i++{
		var art []Models.Article
		Mysql.DB.Where("article_sort=?",articlesort[i].Id).Find(&art)
		for j:=0;j< len(art);j++{
			article=append(article,art[j])
		}
	}
	return article
}
