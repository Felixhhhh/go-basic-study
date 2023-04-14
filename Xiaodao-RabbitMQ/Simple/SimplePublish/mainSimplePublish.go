package main

import (
	"fmt"
	"rabbit/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("xiaodao")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello xiaodao!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}