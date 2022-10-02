package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func init() {
	log.Println("creating topics...")
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "create_company", 0)
	if err != nil {
		log.Println("creating topic create_company: ", err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             "update_company",
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Println("creating topic update_company: ", err)
	}
	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             "delete_company",
		NumPartitions:     2,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Println("creating topic delete_company: ", err)
	}

	log.Println("created topics...")
}
func EventProducer(body []byte, topic string, id string) {
	msg := kafka.Message{
		Key:   []byte(id),
		Value: body,
	}

	kafkaWriter := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	err := kafkaWriter.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Println("error writing msg: ", err)
	}
}
