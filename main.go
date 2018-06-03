package main

import (
	server "github.com/anovafawzi/socialmedia/server"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	server.LoadRoute().Run()
}
