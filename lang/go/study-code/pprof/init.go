package pprof

import (
	"fmt"
	_ "net/http/pprof"
)

var datas []string

//func main() {
//	go func() {
//		for {
//			log.Println("len: %d", Add("go-programming-tour-book"))
//			time.Sleep(10 * time.Millisecond)
//		}
//	}()
//	_ = http.ListenAndServe("0.0.0.0:6060", nil)
//}
//
//func Add(str string) int {
//	data := []byte(str)
//	datas = append(datas, string(data))
//	return len(datas)
//}
func Testwork() {
	fmt.Println("work")
}
func init() {

	fmt.Println("init pprof")
}
