package schema

import (
	"fmt"
	"reflect"

	"github.com/uudashr/orm/step1/log"
)

//一个Filed就是一张表的一个字段
type Field struct {
	Name string
	Type string
	Tag  string
}

//Schema相当于一张表(Table)的抽象化
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field //便于遍历
}

//获取mysql类型的函数
func DataTypeOf(typ reflect.Value) string {
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

/*
type User struct{
	Name string
	Id int
}
目的是将上述结构体解析成Schema
*/
func Parse(dest interface{}) *Schema {
	//首先传的是一个指针
	//本句是获取结构体对象的实例，返回的是结构体指针所指的具体对象的reflect.Value
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type() //其本质其实和Elem()方法一样
	s := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i) //获取第i个字段的全部信息,p.name
		field := &Field{
			Name: p.Name,
			Type: DataTypeOf(reflect.Indirect(reflect.New(p.Type))), //reflect.New(type方法)获取一个value
		}
		log.Infof("第%d个解析的行", i)
		log.Info(field.Name, field.Type)
		//暂且不考虑tag标签
		s.Fields = append(s.Fields, field)
		s.FieldNames = append(s.FieldNames, p.Name)
		s.fieldMap[p.Name] = field
	}
	return s
}
