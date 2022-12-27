package xnsq

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func StartNsq() {
	go startConsumer("consume1")
	go startConsumer("consume2")
	startProducer()
}

var url string

func init() {
	//具体ip,端口根据实际情况传入或者修改默认配置
	flag.StringVar(&url, "url", "0.0.0.0:58066", "nsqd")
	flag.Parse()
}

var produceCnt, consumer1Cnt, consumer2Cnt int = 0, 0, 0

// 生产者
func startProducer() {
	cfg := nsq.NewConfig()
	fmt.Println("url is:", url)
	producer, err := nsq.NewProducer(url, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 发布消息
	for {
		if err := producer.Publish("test", []byte("test message")); err != nil {
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
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		switch consumerName {
		case "consume1":
			log.Println(consumerName+"收到的消息"+string(message.Body), consumer1Cnt)
			consumer1Cnt++
		case "consume2":
			log.Println(consumerName+"收到的消息"+string(message.Body), consumer2Cnt)
			consumer2Cnt++
		}
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD(url); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
	fmt.Println("consumer.StopChan")
}
