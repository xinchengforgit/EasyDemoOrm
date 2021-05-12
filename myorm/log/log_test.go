package log

import (
	"myorm/log"
	"testing"
)

//测试手写的log库
func Test_log(t *testing.T) {
	log.Error("hello")
	log.Info("hello")
}
