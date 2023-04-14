package main

import "rabbit/RabbitMQ"

func main() {
	kutengone := RabbitMQ.NewRabbitMQRouting("xiaodao", "xiaodao_one")
	kutengone.RecieveRouting()
}
