package Dao

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Models"
)

func InsertArticleMapper(article Models.Article)(err error)  {
	//Mysql.DB.Raw("insert into tb_article (title,content,created_time) values('test','testcontent','2020-08-28 11:33:28')")
	dberr:=Mysql.DB.Create(&article)
	err=dberr.Error
	return
}
