package client

import (
	"fmt"
	"github.com/MahmoudMekki/XM-Task/config"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

var producer sarama.SyncProducer
var consumer sarama.Consumer

func establishKafkaProducer() {
	var err error
	prodConfig := sarama.NewConfig()
	prodConfig.Producer.Return.Successes = true
	url := fmt.Sprintf("%s:%s", config.GetEnvVar("BROKER_HOST"), config.GetEnvVar("BROKER_PORT"))
	producer, err = sarama.NewSyncProducer([]string{url}, prodConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to establish connection with Kafka")
	}
}
func establishKafkaConsumer() {
	var err error
	consumerConfig := sarama.NewConfig()
	url := fmt.Sprintf("%s:%s", config.GetEnvVar("BROKER_HOST"), config.GetEnvVar("BROKER_PORT"))
	log.Info().Msg(url)
	consumer, err = sarama.NewConsumer([]string{url}, consumerConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to establish connection with Kafka")
	}
}
func GetProducer() sarama.SyncProducer {
	if producer == nil {
		establishKafkaProducer()
	}
	return producer
}
func GetConsumer() sarama.Consumer {
	if consumer == nil {
		establishKafkaConsumer()
	}
	return consumer
}
