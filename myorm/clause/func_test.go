package clause

import (
	"fmt"
	"strings"
	"testing"
)

type User struct {
	name string
	age  int
}

//Test***开头的函数
func Test_func(t *testing.T) {
	str := genBindVars(5)
	fmt.Println(str)
	//?, ?, ?, ?, ?结果如下,就是生成了占位符
}

func Test_v(t *testing.T) {
	a := &User{"leh", 18}
	fmt.Printf("%v", a)
	fmt.Printf("%v\n", *a)
	var b []string
	b = append(b, "hello")
	b = append(b, "world")
	c := strings.Join(b, ",")
	fmt.Printf("%v\n", b) //[hello world]
	fmt.Printf("%v\n", c) //hello,world
	//结果如下&{leh 18}{leh 18}
}
