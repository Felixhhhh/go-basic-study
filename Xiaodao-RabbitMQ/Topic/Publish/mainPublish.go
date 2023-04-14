package main

import (
	"fmt"
	"rabbit/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	kutengOne := RabbitMQ.NewRabbitMQTopic("exXiaodaoTopic", "xiaodao.topic.one")
	kutengTwo := RabbitMQ.NewRabbitMQTopic("exXiaodaoTopic", "xiaodao.topic.two")
	for i := 0; i <= 100; i++ {
		kutengOne.PublishTopic("Hello xiaodao topic one!" + strconv.Itoa(i))
		kutengTwo.PublishTopic("Hello xiaodao topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}