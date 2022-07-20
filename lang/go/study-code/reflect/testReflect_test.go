package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type TestReflect struct {
	a string
}

func (U *TestReflect) HelloTestReflectFunc() {
	fmt.Println("hello_test_reflect")
}

func (U *TestReflect) HalloTestReflectFunc() {
	fmt.Println("hello_test_reflect2")
}

func getMethod(group interface{}) {
	groupType := reflect.TypeOf(group)
	//groupName := groupType.Elem().Name()
	groupValue := reflect.ValueOf(group)
	fmt.Println(groupType.Name(), groupValue, groupType.NumMethod())
	for i := 0; i < groupType.NumMethod(); i++ {
		m := groupType.Method(i)
		name := m.Name
		firstCh := name[:1]
		fmt.Println(name, groupValue, firstCh)
		m.Func.Call([]reflect.Value{groupValue})

	}
}

func Test_ReflectWrok(t *testing.T) {

	group := TestReflect{"name"}
	var testCases []interface{}
	testCases = append(testCases, &group)
	getMethod(testCases[0])
	fmt.Println("hello!")

	val := []string{}
	res := reflect.TypeOf(val)
	fmt.Println(res, reflect.TypeOf(val).String())

}
