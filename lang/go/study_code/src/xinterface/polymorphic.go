package xinterface

import "fmt"

type notifier interface {
	notify()
	changeValue()
	varChangeValue()
}

type user struct {
	notifier
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("user.name: ", u.name)
}

func (u *user) changeValue() {
	u.name = u.name + u.name
}

func (u user) varChangeValue() {
	u.email = u.email + u.email
}

type vipUser struct {
	user
	vipScore int
}

func (vu *vipUser) changeValue() {
	vu.name = vu.name + vu.name
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("%p admin notify: %v %v ", a, a.name, a.email)
}

func (a *admin) changeValue() {
	a.name = a.name + a.name
}

func (a admin) varChangeValue() {
	a.email = a.email + a.email
}

func sendNotification(n notifier) {
	n.notify()
}

func changeSelfValue(n notifier) {
	n.changeValue()
}

func varChangeSelfValue(n notifier) {
	n.varChangeValue()
}

func mainWork() {
	bill := user{name: "bill", email: "1234@qq.com"}
	fmt.Printf("%p\n", (&bill))
	fmt.Printf("user.p %p user.notifier %+v user.notifier.p %p\n", &bill, bill.notifier, &(bill.notifier))
	//sendNotification(&bill)
	changeSelfValue(&bill)
	fmt.Println("change bill.name use pointer", bill)

	lisa := admin{name: "lisa", email: "7890@qq.com"}
	fmt.Printf("%p\n", &lisa)
	sendNotification(&lisa)
	varChangeSelfValue(&lisa)
	fmt.Println("change lisa.name use var", lisa)

	bob := vipUser{
		user:     user{name: "bob", email: "6666@qq.com"},
		vipScore: 0,
	}
	bob.notify()
}
