package xruntime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func Test_AssertionWrok(t *testing.T) {
	asrt := assert.New(&testing.T{})

	res := asrt.Equal(2, 1)
	fmt.Println(res)

	res1 := asrt.Equal(1, 1)
	fmt.Println(res1)

	fmt.Println(runFuncName())
}
