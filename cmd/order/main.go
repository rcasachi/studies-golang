package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcasachi/studies-golang/internal/infra/database"
	"github.com/rcasachi/studies-golang/internal/usecases"
	"github.com/rcasachi/studies-golang/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	usecase := usecases.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // T2

	rabbitmqWorker(msgRabbitmqChannel, usecase)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecases.CalculateFinalPrice) {
	fmt.Println("Starting RabbitMQ")

	for msg := range msgChan {
		var input usecases.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Message processing and saved at database:", output)
	}
}
