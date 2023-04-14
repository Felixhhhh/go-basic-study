package main

import "rabbit/RabbitMQ"

func main() {
	xiaodaoTwo := RabbitMQ.NewRabbitMQTopic("exXiaodaoTopic", "xiaodao.*.two")
	xiaodaoTwo.RecieveTopic()
}
