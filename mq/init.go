package mq

import (
	"chat/conf"
	"strings"

	"github.com/streadway/amqp"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	connString := strings.Join([]string{conf.RabbitMQ, "://", conf.RabbitMQUser, ":", conf.RabbitMQPassWord, "@", conf.RabbitMQHost, ":", conf.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	RabbitMq = conn
}
