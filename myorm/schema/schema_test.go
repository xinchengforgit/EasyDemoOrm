package schema

import (
	"myorm/log"
	"testing"
)

type User struct {
	Name string
	Age  int
}

//用于测试解析model

func Test_parse_schema(t *testing.T) {
	schema := Parse(&User{})
	log.Info(schema.Name)
	for _, name := range schema.FieldNames {
		log.Info(name)
		field := schema.fieldMap[name]
		log.Info(field.Name, field.Type)
	}
}
func Test_record(t *testing.T) {

}
