package models

import (
	"encoding/json"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var (
	//Topic : 
	Topic = "imgNotification"
)
type kafkaM struct {
	Image string	`json:"image"`
	Album string	`json:"album"`
	Operation string	`json:"operation"`
}

//NewKafkaMessage :
func NewKafkaMessage(image, album, op string) *kafka.Message{
	msg := &kafkaM{Image: image,Album: album,Operation: op}
	val, _ := json.Marshal(msg)
	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &Topic,Partition: kafka.PartitionAny},
		Value: val,
	}
}