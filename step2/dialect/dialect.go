package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    //用于将Go的类型转换为该数据库的类型
	TableExistSQL(tableName string) (string, []interface{}) //返回某个表是否存在的 SQL 语句，参数是表名(table)
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
