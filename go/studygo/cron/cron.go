package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

/*
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

https://pkg.go.dev/github.com/robfig/cron
*/

var cronTab = cron.New()
var spec = "* * * * * *"
var spec1 = "5-40/8 * * * * *" // 第5秒-第40秒区间，每隔8秒种打印一次
var spec2 = "4/5 * * * * *"    // 第4秒开始-第59秒，每隔5秒打印一次(4/5 = 4-59/5)
var spec3 = "0 0 23 * * 0,1,2" // 每周一，周二，周五的晚上11点整打印一次
var spec4 = "*/9 29 0 1 * *"   // 每个月的一号，0点第29分钟，每隔9秒打印一次
var spec5 = "0 30 * * * *"

func cronWork() {
	cronTab.AddFunc(spec, func() {
		fmt.Println(time.Now())
	})
	//cronTab.Start()

	cronTab.Run()

}

func scheduleWork(thisSpec string, nextN int) {
	schedule, err := cron.Parse(thisSpec)
	if err != nil {
		fmt.Printf("scheduleWork err(%s)", err)
		return
	}

	now := time.Now()

	fmt.Println("now time is:", now)

	for nextN > 0 {
		nextN--

		next := schedule.Next(now)

		str := fmt.Sprintf("next time is:%v\n", next)
		if nextN == 0 {
			str += "\n"
		}

		fmt.Printf(str)

		now = next
	}

}

func main() {

	//scheduleWork(spec, 5)
	//scheduleWork(spec1, 10)
	//scheduleWork(spec2, 10)
	//scheduleWork(spec3, 10)

	scheduleWork(spec5, 20)

	<-time.After(time.Second * 10)
}
