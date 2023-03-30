package alga

import (
	"fmt"
	"log"
	"os"
)

var LogLevel int = -100

var LogLevelN1 = -1
var LogLevelN100 = -100

//func log(s string, i ...interface{}) {
//	//fmt.Printf(s, i...)
//}

func llog(ll int, s string, i ...interface{}) {
	if ll <= LogLevel {
		fmt.Printf(s, i...)
	}
}

func init() {
	file := "./out.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[case]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
