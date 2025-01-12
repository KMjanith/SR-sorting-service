package main

import (
	"context"
	"log"
	"os"
	"sorting/consumers"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func main() {

	godotenv.Load() //make the connection to the mongodb
	log.Println("starting sorting service")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Set up a context with timeout
	defer cancel()

	//---------------------------------------RABBITMQ CONNECTION---------------------------------//

	rabbitUrl := os.Getenv("RABBITMQ_URI")
	conn, err := amqp.Dial(rabbitUrl) //making the connection to the rabbitmq
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	ctx = context.WithValue(ctx, "conn", conn) //adding the connection to the context
	consumerChannel, err := conn.Channel()     //Open a channel for consumer
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer consumerChannel.Close()
	ctx = context.WithValue(ctx, "consumerChannel", consumerChannel) //adding the consumer channel to the context
	producerChannel, err := conn.Channel()                           //open a channel for producer
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer producerChannel.Close()
	log.Println("Connected to RabbitMQ")
	ctx = context.WithValue(ctx, "producerChannel", producerChannel) //adding the producer channel to the context

	consumers.Consumer(ctx) //run the consumer
}
