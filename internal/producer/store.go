package producer

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-register/internal/producer/model"
)

// StoreSuperhero publishes new Superhero on Kafka topic for it to be
// consumed by record keeper and stored in DB and Elasticsearch.
func(p *Producer) StoreSuperhero(s model.Superhero) error {
	var sb bytes.Buffer

	key := s.ID

	err := json.NewEncoder(&sb).Encode(s)
	if err != nil {
		return err
	}

	err = p.Producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: sb.Bytes(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
