package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 消息处理函数
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("收到消息: %s\n", msg.Payload())
	fmt.Printf("主题: %s\n\n", msg.Topic())
}

func main() {
	// MQTT 配置
	broker := "tcp://localhost:1883" // 修改为你的 MQTT Broker 地址
	clientID := "go-mqtt-subscriber"
	topic := "test/topic"

	// 设置 MQTT 客户端选项
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("连接断开: %v", err)
	})

	// 创建并连接客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("连接错误: %v", token.Error())
	}
	fmt.Println("成功连接到 MQTT Broker")

	// 订阅主题
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("订阅错误: %v", token.Error())
	}
	fmt.Printf("已订阅主题: %s\n", topic)

	// 等待退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// 断开连接
	client.Unsubscribe(topic)
	client.Disconnect(250)
	fmt.Println("已断开连接")
}
