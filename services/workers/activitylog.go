package main

import (
	"encoding/json"
	"fmt"
	"github.com/MahmoudMekki/XM-Task/client"
	"github.com/MahmoudMekki/XM-Task/database"
	"github.com/MahmoudMekki/XM-Task/pkg/models"
	"github.com/MahmoudMekki/XM-Task/pkg/repo/logsDAL"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := database.CreateDBConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func main() {
	worker := client.GetConsumer()
	consumer, err := worker.ConsumePartition(models.ActivityLogTopic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start consumer for topic " + models.ActivityLogTopic)
	}
	log.Info().Msg("Successfully started consumer for topic " + models.ActivityLogTopic)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Err(err).Msg("Error from consumer")
			case msg := <-consumer.Messages():
				var activity models.Log
				json.Unmarshal(msg.Value, &activity)
				err := logsDAL.CreateLog(activity)
				if err != nil {
					log.Err(err).Msg("Failed to create log")
				}
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	log.Info().Msg("Closing consumer")
	if err := worker.Close(); err != nil {
		panic(err)
	}
}
