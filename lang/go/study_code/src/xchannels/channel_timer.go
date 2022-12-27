package xchannels

import (
	"fmt"
	"time"
)

func timeAfter(str string) {
	fmt.Println(str + "timeAfter before")
	<-time.After(time.Second * 3)
	fmt.Println(str + "timeAfter after")
	fmt.Println(time.Now())

}

func invokeTimeAfter(str string) {
	timeAfter(str)
}
