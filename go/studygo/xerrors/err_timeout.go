package xerrors

import (
	"fmt"
	"net"
)

type nerr struct {
	timeout int
}

func (e nerr) Timeout() bool {
	return true
}
func (e nerr) Error() string {
	return "err"
}

func (e nerr) Temporary() bool {
	return true
}

func returnERR() error {
	var ne nerr
	return &ne
}

func defineErrIsTimeout() {
	ne := returnERR()
	err, ok := ne.(net.Error)

	fmt.Println(ok, err.Timeout())

	if ok {
		if err.Timeout() {
			fmt.Println("this is a timeout err")
		}
	}
}
