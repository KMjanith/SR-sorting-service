package producers

import (
	"context"
	"log"
	"sorting/spec"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func SendSortedREsponse(method string, sortedArr []int64, time string, ctx context.Context, ReplyTo string) {

	msg := &spec.SoritingResponse{
		Method:    method,
		SortedArr: sortedArr,
		Time:      time,
	}

	log.Println(sortedArr)

	ch := ctx.Value("producerChannel")

	// Serialize message using protobuf
	request, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.(*amqp.Channel).PublishWithContext(ctx,
		"",      // exchange
		ReplyTo, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        request,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
	log.Printf(" [x] Sent %s\n", msg)
}
