package Mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var DB *gorm.DB

func init()  {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/omdevice?charset=utf8&parseTime=True&loc=Local")
	//DB, err = gorm.Open("mysql", "root:134120@tcp(127.0.0.1:3306)/iscm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}