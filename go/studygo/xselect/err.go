package xselect

import (
	"fmt"
	"testing"
	"time"
)

func Test_ZpTimer_Set2(t1 *testing.T) {

	stop := time.After(100 * time.Second)

	//second2 := time.After(time.Second*2)
	second4 := time.After(time.Second * 4)
	second6 := time.After(time.Second * 6)

	for {
		select {
		case <-second6:
			fmt.Println("xxxafter 6 second")
		case <-second4:
			fmt.Println("xxafter 4 second")
		case <-time.After(time.Second * 2): //！实际上效果类似于tick，原因是每次for的时候time.After都会重新生成一个2秒后的定时器
			fmt.Println("xafter 2 second")
		case <-stop:
			fmt.Println("now we return")
			return
		}
	}
}
func Test_ZpTimer_Set3(t1 *testing.T) {

	stop := time.After(100 * time.Second)

	//	下面这段程序如果不是stop，则永远都不会停止
	//  会一直打印xafter 2 second
	for {
		select {
		case <-time.After(time.Second * 6):
			fmt.Println("xxxafter 6 second")
		case <-time.After(time.Second * 4):
			fmt.Println("xxafter 4 second")
		case <-time.After(time.Second * 2): //！实际上效果类似于tick，原因是每次for的时候time.After都会重新生成一个2秒后的定时器
			fmt.Println("xafter 2 second")
		case <-stop:
			fmt.Println("now we return")
			return
		}
	}
}
