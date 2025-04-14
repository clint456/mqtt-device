package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// MQTT 配置
	broker := "tcp://localhost:1883" // 修改为你的 MQTT Broker 地址
	clientID := "go-mqtt-publisher"
	topic := "test/topic"

	// 设置 MQTT 客户端选项
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("连接断开: %v", err)
	})

	// 创建并连接客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("连接错误: %v", token.Error())
	}
	fmt.Println("成功连接到 MQTT Broker")

	// 发布消息
	ticker := time.NewTicker(2 * time.Second) // 每2秒发布一次
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	i := 0
	for {
		select {
		case <-ticker.C:
			message := fmt.Sprintf("消息 #%d from Go publisher", i)
			token := client.Publish(topic, 0, false, message)
			token.Wait()
			if token.Error() != nil {
				log.Printf("发布消息错误: %v", token.Error())
			} else {
				fmt.Printf("已发布消息: %s\n", message)
			}
			i++
		case <-sigChan:
			ticker.Stop()
			client.Disconnect(250)
			fmt.Println("已断开连接")
			return
		}
	}
}
