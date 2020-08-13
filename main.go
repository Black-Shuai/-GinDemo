package main

import (
	Mysql "GinDemo/Databases"
	"GinDemo/Router"
)

func main() {
	defer Mysql.DB.Close()
	Router.InitRouter()
}