package notification

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

//Client :
type Client interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}

//ConnectKafka :
func ConnectKafka(host string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Errorf("Error connecting to Kafka %v", err)
		return nil
	}
	return p
}
