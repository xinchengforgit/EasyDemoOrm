package session

import (
	"database/sql"
	"myorm/log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_session(t *testing.T) {
	log.Info("hello")
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Error(err)
	}
	s := New(db)
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec() //Ok这是基本的操作
}
