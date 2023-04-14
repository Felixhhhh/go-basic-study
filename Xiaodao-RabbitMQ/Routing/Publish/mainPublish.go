package main

import (
	"fmt"
	"rabbit/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	kutengone := RabbitMQ.NewRabbitMQRouting("xiaodao", "xiaodao_one")
	kutengtwo := RabbitMQ.NewRabbitMQRouting("xiaodao", "xiaodao_two")
	for i := 0; i <= 100; i++ {
		kutengone.PublishRouting("Hello xiaodao one!" + strconv.Itoa(i))
		kutengtwo.PublishRouting("Hello xiaodao Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}