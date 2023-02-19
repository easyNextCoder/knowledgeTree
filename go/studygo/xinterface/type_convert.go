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
	var sd S
	var sdp *S

	//var cat Cat = RedCat{} //值拷贝
	var cat Cat = &RedCat{} //pointer拷贝
	var catP *Cat = &nilCat
	var redCat RedCat
	var redCatP *RedCat = &RedCat{}

	fmt.Println("------------------start")

	//playWithCat(i)
	//playWithFourFoot(i)
	playWithInterface(i)

	//playWithCat(&i)
	//playWithFourFoot(&i)
	playWithInterface(&i)

	playWithCat(nilCat)
	//playWithFourFoot(nilCat)
	playWithInterface(nilCat)

	playWithCat(cat)
	//playWithFourFoot(cat)
	playWithInterface(cat)

	//playWithCat(catP)
	//playWithFourFoot(catP)
	playWithInterface(catP)

	//playWithCat(&cat)
	//playWithFourFoot(&cat)
	playWithInterface(&cat)

	playWithCat(redCat)
	playWithFourFoot(redCat)
	playWithInterface(redCat)

	playWithCat(&redCat)
	playWithFourFoot(&redCat)
	playWithInterface(&redCat) //&redCat即是*redCat类型也是Cat型
	v := &redCat
	playWithInterface(&v) //15

	playWithCat(redCatP)
	playWithFourFoot(redCatP)
	playWithInterface(redCatP)

	//playWithCat(S{})
	//playWithFourFoot(S{})
	playWithInterface(S{})

	//playWithCat(sd)
	//playWithFourFoot(sd)
	playWithInterface(sd)

	//playWithCat(&sd)
	//playWithFourFoot(&sd)
	playWithInterface(&sd)

	//playWithCat(sdp)
	//playWithFourFoot(sdp)
	playWithInterface(sdp)

	//playWithCat(&sdp)
	//playWithFourFoot(&sdp)
	playWithInterface(&sdp)

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
