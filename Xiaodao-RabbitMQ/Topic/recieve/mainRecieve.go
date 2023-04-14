package main

import "rabbit/RabbitMQ"

func main() {
	xiaodaoOne := RabbitMQ.NewRabbitMQTopic("exXiaodaoTopic", "#")
	xiaodaoOne.RecieveTopic()
}
