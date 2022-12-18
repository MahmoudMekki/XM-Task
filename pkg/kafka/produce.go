package kafka

import (
	"github.com/MahmoudMekki/XM-Task/client"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

func Produce(topic string, obj []byte) {
	producer := client.GetProducer()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(obj),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Err(err).Msg("Failed to send message to Kafka")
	}
	log.Info().Msgf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)
}
