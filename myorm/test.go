package main

import (
	"database/sql"
	"fmt"
	"myorm/session"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		panic(err)
	}

	s := session.New(db)
	s.Model(&User{})
	s.DropTable()   //如果存在就先删除
	s.CreateTable() //以table来创建表
	//ok
	user1 := &User{"leh", 18}
	user2 := &User{"zhangsan", 15}
	s.Insert(user1, user2)
	var users []User
	s.Find(&users)
	for _, usert := range users {
		fmt.Println(usert)
	}
	_, err1 := s.Where("Name = ?", "zhangsan").Update("Age", 30)
	if err1 != nil {
		panic(err1)
	}
	s.Where("Name=?", "zhangsan").Delete()
}
