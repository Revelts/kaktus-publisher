package main

import (
	"kaktus-services/Connections"
	"kaktus-services/Routes"
	"log"
)

func main() {
	var err error
	err = Connections.InitiateConnection()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}
	err = Routes.HandleRequests()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	channel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}
	err = channel.ExchangeDeclare("kaktus", "topic",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}
}
