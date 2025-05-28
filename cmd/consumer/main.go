package main

import (
	"car-service/internal/model"
	"context"
	"encoding/json"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
		Topic:   "vehicle.created",
		GroupID: "vehicle-consumer-group",
	})

	for {
		var v model.Vehicle
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Errorf("failed to read message: %v", err)
			continue
		}
		err = json.Unmarshal(msg.Value, &v)
		if err != nil {
			logrus.Errorf("failed to unmarshal message: %v", err)
			continue
		}
		logrus.Infof("vehicle created: id=%d, make=%s, mark=%s, year=%d", v.ID, v.Make, v.Mark, v.Year)
	}
}
