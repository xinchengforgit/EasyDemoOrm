package dialect

import (
	"fmt"
	"reflect"
)

type Mysql struct{}

func init() {
	RegisterDialect("mysql", &Mysql{})
}

//返回表的数据
func (s *Mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "text"
	case reflect.Float32, reflect.Float64:
		return "float"
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

//判断要建的表是否在数据库中存在
func (s *Mysql) TableExistSQL(tableName string, dbName string) (string, []interface{}) {
	args := []interface{}{tableName, dbName}
	//该语句用来查询该数据库中是否存在该表
	return " select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA=? and TABLE_NAME=?;", args
}
