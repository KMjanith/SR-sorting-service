package consumers

import (
	"context"
	"log"
	"sorting/service"
	"sorting/spec"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

// Consumer1 is a simple consumer that receives a message from the queue
func Consumer(ctx context.Context) {

	CONSUMER_CHANNEL := ctx.Value("consumerChannel").(*amqp.Channel)

	ApiSortmsg, err := CONSUMER_CHANNEL.QueueDeclare(
		"ApiSortMsg", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare Api-Auth-msg queue: %v", err)
	}

	msgs, err := CONSUMER_CHANNEL.Consume(
		ApiSortmsg.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process messages
	forever := make(chan bool)

	go func() {
		for d := range msgs {

			var SortingRequest spec.SortingRequest
			// Unmarshal the Protobuf message
			err := proto.Unmarshal(d.Body, &SortingRequest)
			if err != nil {
				log.Fatalf("Failed to unmarshal Protobuf message: %v", err)
			}
			// Accessing fields of the protobuf message
			log.Printf("Received a SortingRequest: method=%s", SortingRequest.Method)
			log.Println("calling sorting")
			method := SortingRequest.Method
			arr := SortingRequest.Arr

			switch method {
			case "bubble":
				service.BubbleSort(arr, ctx, d.ReplyTo)
			case "quick":
				low := 0
				high := len(arr) - 1
				log.Println(d.ReplyTo)
				service.QuickSort(arr, low, high, ctx,d.ReplyTo)
			case "merge":
				left := 0
				right := len(arr) - 1
				service.MergeSort(arr, left, right, ctx,d.ReplyTo)
			case "insertion":
				service.InsertionSort(arr, ctx,d.ReplyTo)
			case "select":
				service.SelctionSort(arr, ctx,d.ReplyTo)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
