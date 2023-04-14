package main

import "rabbit/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("xiaodao")
	rabbitmq.ConsumeSimple()

}