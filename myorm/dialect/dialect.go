package dialect

import "reflect"

var DialectMap = map[string]Dialect{} //接口的声明是这样的

type Dialect interface {
	DataTypeOf(typ reflect.Value) string //判断表是否存在
	TableExistSQL(tableName string, dbname string) (string, []interface{})
}

func RegisteDialect(name string, new_dialect Dialect) {
	DialectMap[name] = new_dialect
	return
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = DialectMap[name] //注意map获取的值返回有两个,后者是bool类型
	return
}
