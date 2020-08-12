package main

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Router"
	//"Gindemo/Databases"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}