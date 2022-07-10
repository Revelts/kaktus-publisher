package Connections

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var RabbitConnection *amqp.Connection

func rabbitConn() (con *amqp.Connection, err error) {
	con, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	return
}

func InitiateConnection() (err error) {
	RabbitConnection, err = rabbitConn()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}
	return
}
