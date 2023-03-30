package xinterface

import (
	"fmt"
	"reflect"
)

type Cat interface {
	Miao() string
}

type FourFoot interface {
	Work() string
}

type RedCat struct {
	i int
}

func (r RedCat) Miao() string {
	return "RedCat Miao"
}

func (r RedCat) Work() string {
	return "RedCat Work"
}

type RedBigCat struct {
	RedCat
}

type BlueCat struct {
	v int
}

func (r BlueCat) Miao() string {
	return "BlueCat Miao"
}

func playWithCat(cat Cat) {
	//cat.Miao()
	switch cat.(type) {
	case Cat:
		fmt.Println("type is Cat", cat.(FourFoot).Work())
	case FourFoot:
		fmt.Println("type is FourFoot")
	case RedCat:
		fmt.Println("type is RedCat")
	//case *Cat:
	//	fmt.Println("type is *Cat")
	case *RedCat:
		fmt.Println("type is *RedCat", cat.(Cat).Miao())
	case interface{}:
		fmt.Println("type is interface")
	case nil:
		fmt.Println("type is nil")
	default:
		fmt.Println("enter default")
	}
}

func playWithFourFoot(cat FourFoot) {
	//cat.Miao()
	switch cat.(type) {
	case Cat:
		fmt.Println("type is Cat", cat.(FourFoot).Work())
	case FourFoot:
		fmt.Println("type is FourFoot")
	case RedCat:
		fmt.Println("type is RedCat")
	//case *Cat:
	//	fmt.Println("type is *Cat")
	case *RedCat:
		fmt.Println("type is *RedCat", cat.(Cat).Miao())
	case interface{}:
		fmt.Println("type is interface")
	case nil:
		fmt.Println("type is nil")
	default:
		fmt.Println("enter default")
	}
}

func playWithInterface(cat interface{}) {
	//cat.Miao()
	switch cat.(type) {

	case *Cat:
		fmt.Println("type is *Cat")
	case *RedCat:
		fmt.Println("type is *RedCat", cat.(Cat).Miao())
	case Cat:
		fmt.Println("type is Cat", cat.(FourFoot).Work())
	case FourFoot:
		fmt.Println("type is FourFoot")
	case RedCat:
		fmt.Println("type is RedCat")
	case **RedCat:
		fmt.Println("type is **RedCat", (**(cat.(**RedCat))).Miao())
	case nil:
		fmt.Println("type is nil")
	case S:
		fmt.Println("type is S")
	case *S:
		fmt.Println("type is *S")
	case **S:
		fmt.Println("type is **S")
	case *interface{}:
		fmt.Println("type is *interface{}")
	case interface{}:
		fmt.Println("type is *interface{}")
	default:
		fmt.Println("enter default")
	}
}

type S struct {
	int
	float64
	i []int
}

func typeConvertWork() {

	var i interface{} = S{}
	var nilCat Cat

	//var cat Cat = RedCat{} //值拷贝
	var cat Cat = &RedCat{} //pointer拷贝
	var catP *Cat = &nilCat
	var redCat RedCat
	var redCatP *RedCat = &RedCat{}

	fmt.Println("------------------start")

	//playWithCat(i)
	//playWithFourFoot(i)
	playWithInterface(i) //type S

	//playWithCat(&i)
	//playWithFourFoot(&i)
	playWithInterface(&i) //type *interface{} (取地址操作再赋值给interface{}，会覆盖之前interface中存的类型)

	playWithCat(nil)       //type nil
	playWithFourFoot(nil)  //type nil
	playWithInterface(nil) //type nil

	playWithCat(nilCat) //type nil
	//playWithFourFoot(nilCat)
	playWithInterface(nilCat) //type nil

	playWithCat(cat) //type Cat
	//playWithFourFoot(cat)
	playWithInterface(cat) //type *RedCat

	//playWithCat(catP)
	//playWithFourFoot(catP)
	playWithInterface(catP) //input type &(Cat类型的空接口)  type *Cat

	//playWithCat(&cat)
	//playWithFourFoot(&cat)
	playWithInterface(&cat) //type *Cat (取地址操作再赋值给interface{}，会覆盖之前interface中存的类型)

	playWithCat(redCat)       //type Cat
	playWithFourFoot(redCat)  //type Cat
	playWithInterface(redCat) //type Cat

	playWithCat(&redCat)       //自身被赋值给了函数的参数，所以在内部的判断类型是redCat
	playWithFourFoot(&redCat)  //自身被赋值给了函数的参数，所以在内部的判断类型是redCat
	playWithInterface(&redCat) //&redCat即是*redCat类型也是Cat型
	v := &redCat
	playWithInterface(&v) //type **RedCat

	playWithCat(redCatP)       //Cat
	playWithFourFoot(redCatP)  //type Cat
	playWithInterface(redCatP) //type *RedCat

	var sd S
	var sdp *S
	//playWithCat(S{})
	//playWithFourFoot(S{})
	playWithInterface(S{}) //type S

	//playWithCat(sd)
	//playWithFourFoot(sd)
	playWithInterface(sd) //type S

	//playWithCat(&sd)
	//playWithFourFoot(&sd)
	playWithInterface(&sd) //type *S

	//playWithCat(sdp)
	//playWithFourFoot(sdp)
	playWithInterface(sdp) //type *S

	//playWithCat(&sdp)
	//playWithFourFoot(&sdp)
	playWithInterface(&sdp) //type **S

	fmt.Println("------------------end")

}

func typeConvertPanic() {
	var b interface{}
	b = &BlueCat{v: 1}
	cat := b.(FourFoot)
	fmt.Println(cat.Work())
}

func stringSliceConvert() {
	var ss []byte
	var i interface{} = &ss
	switch i.(type) {
	case *[]interface{}:
		fmt.Println("*[]interface{}")
	case []string:
		fmt.Println("[]string")
	case *string:
		fmt.Println("*string")
	//case *[]string:
	//	fmt.Println("*[]string")
	default:
		//func convertAssign(d interface{}, s interface{}) (err error) {redigo/redis/scan.go对于这段有详细的用法代码
		if d := reflect.ValueOf(i); d.Type().Kind() == reflect.Ptr {
			innerV := d.Elem()
			fmt.Println(innerV.Type().Kind())
			innerV.SetBytes([]byte(string("[123]")))
		}
		fmt.Println("")
	}
	fmt.Println(string(ss))
}
