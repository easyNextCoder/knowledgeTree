package xinterface

import (
	"fmt"
	"strings"
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

func typeFinder(i interface{}) string {

	var types []string = []string{"typeFinder->"}

	_, ok0 := i.(S)
	if ok0 {
		types = append(types, "S")
	}

	_, ok_01 := i.(*S)
	if ok_01 {
		types = append(types, "*S")
	}

	_, ok := i.(Cat)
	if ok {
		types = append(types, "Cat")
	}

	_, ok2 := i.(*Cat)
	if ok2 {
		types = append(types, "*Cat")
	}

	_, ok3 := i.(FourFoot)
	if ok3 {
		types = append(types, "FourFoot")
	}

	_, ok4 := i.(*FourFoot)
	if ok4 {
		types = append(types, "*FourFoot")
	}

	_, ok5 := i.(RedCat)
	if ok5 {
		types = append(types, "RedCat")
	}

	_, ok6 := i.(*RedCat)
	if ok6 {
		types = append(types, "*RedCat")
	}

	_, ok7 := i.(RedBigCat)
	if ok7 {
		types = append(types, "RedBigCat")
	}

	_, ok8 := i.(*RedBigCat)
	if ok8 {
		types = append(types, "*RedBigCat")
	}

	_, ok9 := i.(interface{})
	if ok9 {
		types = append(types, "interface{}")
	}

	_, ok10 := i.(*interface{})
	if ok10 {
		types = append(types, "*interface{}")
	}

	return strings.Join(types, ",") + "\n"
}

func playWithCat(cat Cat) {
	//cat.Miao()
	switch cat.(type) {
	//case *Cat://error
	//	fmt.Println("type is *Cat")
	case Cat:
		fmt.Println("type is Cat") // cat.(FourFoot).Work())
	case FourFoot:
		fmt.Println("type is FourFoot")
	case RedCat:
		fmt.Println("type is RedCat")
	//case *Cat:
	//	fmt.Println("type is *Cat")
	case *RedCat:
		fmt.Println("type is *RedCat") //, cat.(Cat).Miao())
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
	case interface{}:
		fmt.Println("type is interface{}")
	default:
		fmt.Println("enter default")
	}
}

type S struct {
	int
	float64
	i []int
}

func print(a ...interface{}) {
	fmt.Println(a...)
}

func typeConvertWork() {

	//var cat Cat = RedCat{} //值拷贝

	fmt.Println("------------------start")

	var i interface{} = S{}
	fmt.Println("var i interface{} = S{}")
	//playWithCat(i)
	//playWithFourFoot(i)
	playWithInterface(i) //type S
	fmt.Println(typeFinder(i))

	//playWithCat(&i)
	//playWithFourFoot(&i)
	playWithInterface(&i) //type *interface{} (取地址操作再赋值给interface{}，会覆盖之前interface中存的类型)
	fmt.Println(typeFinder(&i))

	playWithCat(nil)       //type nil
	playWithFourFoot(nil)  //type nil
	playWithInterface(nil) //type nil
	fmt.Println(typeFinder(nil))

	var nilCat Cat
	fmt.Println("var nilCat Cat")
	playWithCat(nilCat) //type nil
	//playWithFourFoot(nilCat)
	playWithInterface(nilCat) //type nil
	fmt.Println(typeFinder(nilCat))

	var cat Cat = &RedCat{} //pointer拷贝
	fmt.Println("var cat Cat = &RedCat{}")
	playWithCat(cat) //type Cat
	//playWithFourFoot(cat)
	playWithInterface(cat) //type *RedCat
	print(typeFinder(cat))

	var catP *Cat = &nilCat
	print("var catP *Cat = &nilCat")
	//playWithCat(catP)
	//playWithFourFoot(catP)
	playWithInterface(catP) //input type &(Cat类型的空接口)  type *Cat
	print(typeFinder(catP))

	//playWithCat(&cat)
	//playWithFourFoot(&cat)
	playWithInterface(&cat) //type *Cat (取地址操作再赋值给interface{}，会覆盖之前interface中存的类型)
	print(typeFinder(&cat))

	var redCat RedCat
	print("var redCat RedCat")
	playWithCat(redCat)       //type Cat
	playWithFourFoot(redCat)  //type Cat
	playWithInterface(redCat) //type Cat
	print(typeFinder(redCat))

	playWithCat(&redCat)       //自身被赋值给了函数的参数，所以在内部的判断类型是redCat
	playWithFourFoot(&redCat)  //自身被赋值给了函数的参数，所以在内部的判断类型是redCat
	playWithInterface(&redCat) //&redCat即是*redCat类型也是Cat型
	print(typeFinder(&redCat))
	v := &redCat
	//!!import &&type 不会再有对应的接口了
	playWithInterface(&v) //type **RedCat
	print(typeFinder(&v))

	var redCatP *RedCat = &RedCat{}
	print("var redCatP *RedCat = &RedCat{}")
	playWithCat(redCatP)       //Cat
	playWithFourFoot(redCatP)  //type Cat
	playWithInterface(redCatP) //type *RedCat
	print(typeFinder(redCatP))

	var ii interface{}
	print("var ii interface{}")
	//playWithCat(i) //
	//playWithFourFoot(i)//
	playWithInterface(ii)
	print(typeFinder(ii))

	//ptri := interface{}(nil)
	var ptri *interface{}
	print("var ptri *interface{}")
	//playWithCat(ptri)
	//playWithFourFoot(ptri)//带有方法的interface{} ,只能接受含有此方法的interface{}变量
	playWithInterface(ptri) //空interface{} 可以接受任意类型的变量
	print(typeFinder(ptri))

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
