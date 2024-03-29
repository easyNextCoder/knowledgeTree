package xreflect

import (
	"fmt"
	"reflect"
	"testing"
)

type TestReflect struct {
	a string
	b int
}

func (U *TestReflect) Ha2_aloTestReflectFunc() {
	fmt.Println("hello_test_reflect")
}

func (U *TestReflect) Ha1_lloTestReflectFunc() {
	fmt.Println("hello_test_reflect2")
}

func (U *TestReflect) Ha3_lloTestReflectFunc() {
	fmt.Println("hello_test_reflect2")
}

func getMethod(group interface{}) {

	groupType := reflect.TypeOf(group)
	//groupName := groupType.Elem().Name()
	groupValue := reflect.ValueOf(group)
	fmt.Println(groupType.Name(), groupValue, groupType.NumMethod())
	fmt.Println("groupType.name():", groupType.Elem().Name())

	fmt.Println("groupValue:", groupValue.Elem().Type())
	fmt.Println("groupType.NumMethod():", groupType.NumMethod())

	for i := 0; i < groupType.NumMethod(); i++ { //是按照函数名的字符升序排列
		m := groupType.Method(i)
		name := m.Name
		firstCh := name[:1]
		fmt.Println(name, groupValue, firstCh)
		m.Func.Call([]reflect.Value{groupValue})

	}
}

func Test_ReflectWrok(t *testing.T) {

	group := TestReflect{"name", 10}

	var testCases []interface{}
	testCases = append(testCases, &group)
	getMethod(testCases[0])
	fmt.Println("hello!")

	val := []string{}
	res := reflect.TypeOf(val)
	fmt.Println(res, reflect.TypeOf(val).String())

}
