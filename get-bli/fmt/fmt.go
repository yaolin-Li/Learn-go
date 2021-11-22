package fmt

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("trace.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		//把程序停掉并报错误信息
	}
	Logger = log.New(io.MultiWriter(os.Stdout, file), "Log: ", log.LstdFlags)
}