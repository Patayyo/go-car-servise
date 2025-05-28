package broker

import (
	"car-service/internal/model"
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var writer *kafka.Writer

func InitWriter(topic string) {
	writer = &kafka.Writer{
		Addr:     kafka.TCP("redpanda:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func Publish(msg []byte) error {
	ctx, cansel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cansel()

	logrus.Infof("Writing to Kafka: %s", string(msg))

	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   nil,
		Value: msg,
	})

	if err != nil {
		logrus.Errorf("Kafka write failed: %v", err)
	} else {
		logrus.Infof("Kafka write success!")
	}

	if err != nil {
		logrus.Errorf("failed to publish Kafka message: %v", err)
	}

	return err
}

func PublishVehicleCreated(v model.Vehicle) error {
	logrus.Infof("Publishing vehicle: %+v", v)
	msg, err := json.Marshal(v)
	if err == nil {
		err = Publish(msg)
	}

	return err
}
