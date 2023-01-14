package main

import (
	Godriver "driver"
	"fmt"
)

func main() {
	var db = Godriver.Connect("username", "password", "databasename")
	var data = Godriver.ReturnJSON(db, "select * from user_table")
	fmt.Println(string(data))
}
