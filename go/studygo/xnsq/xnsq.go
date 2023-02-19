package xnsq

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"strconv"
	"time"
)

func NsqServer() {
	args := os.Args
	fmt.Println(args[1])

	if args[1] == "nsq_all" {
		//第一个进程中一个生产者一个消费者
		fmt.Println("xnsq.StartNsq()")
		StartNsq()
	} else {
		//第二个进程中一个消费者
		fmt.Println("xnsq.StartConsumer()")
		StartConsumer()
		time.Sleep(15 * time.Second)
	}
}

func StartNsq() {
	go startConsumer("consume1")
	go startConsumer("consume2")
	startProducer()
}

func StartConsumer() {
	fmt.Println("StartConsumer")
	go startConsumer("consume3")
}

var url string

func init() {
	//具体ip,端口根据实际情况传入或者修改默认配置
	flag.StringVar(&url, "url", "0.0.0.0:58066", "nsqd")
	flag.Parse()
}

var produceCnt, consumer1Cnt, consumer2Cnt, consumer3Cnt int = 0, 0, 0, 0

// 生产者
func startProducer() {
	cfg := nsq.NewConfig()
	fmt.Println("url is:", url)
	producer, err := nsq.NewProducer(url, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 发布消息
	var i int = 0
	for {
		i++
		if err := producer.Publish("test", []byte("(message"+strconv.Itoa(i)+")")); err != nil {
			log.Fatal("publish error: " + err.Error())
		} else {
			fmt.Println("product:", produceCnt)
			produceCnt++
		}
		time.Sleep(1 * time.Second)
	}
}

// 消费者
func startConsumer(consumerName string) {
	fmt.Println("consumer newConsumer start")
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", consumerName, cfg)
	if err != nil {
		fmt.Printf("consumer newConsumer failed")
		log.Fatal(err)
	}
	fmt.Println("consumer newConsumer success")
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		switch consumerName {
		case "consume1":
			log.Println(consumerName+"收到的消息"+string(message.Body), consumer1Cnt)
			consumer1Cnt++
		case "consume2":
			log.Println(consumerName+"收到的消息"+string(message.Body), consumer2Cnt)
			consumer2Cnt++
		case "consume3":
			log.Println(consumerName+"收到的消息"+string(message.Body), consumer3Cnt)
			consumer3Cnt++
		}
		return nil
	}))
	// 连接到单例nsqd
	fmt.Println("url is:", url)
	if err := consumer.ConnectToNSQD(url); err != nil {
		fmt.Printf("consumer 连接错误 %s", err)
		log.Fatal(err)
	}

	fmt.Printf("consumer 连接成功")

	<-consumer.StopChan
	fmt.Println("consumer.StopChan")
}
