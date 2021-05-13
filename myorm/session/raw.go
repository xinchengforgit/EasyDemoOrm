package session

import (
	"database/sql"
	"fmt"
	"myorm/log"
	"myorm/schema"
	"strings"
)

type Session struct {
	db       *sql.DB         //
	sql      strings.Builder //待了解
	sqlVars  []interface{}   //sql语句的参数确实就是这个
	refTable *schema.Schema
}

func New(db *sql.DB) *Session {
	return &Session{
		db: db,
	} //这种返回方式,注意了解一下
}

//每次清除临时的sql语句，增强session的复用性
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

//获取数据库指针
func (s *Session) GetDB() *sql.DB {
	return s.db
}

//Exec是执行当前的sql，并重置session
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.GetDB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

//查询第一个
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.GetDB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.GetDB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ") //向sql字符串中写入数据
	s.sqlVars = append(s.sqlVars, values...)
	fmt.Println(s.sqlVars...)
	return s
}
