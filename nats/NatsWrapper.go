package nats

import (
	"encoding/json"
	"time"

	"github.com/javascript-ninja-123/log-client/domain"
	"github.com/nats-io/stan.go"
)

type INatsWrapper interface {
	Log(name string, message string, level domain.Level) error
	Connect(productionID string, clusterID string, URL string) error
	Subscribe(subject NatsEventType, queueName string, DurableKey string, callback func(data []byte) error)
	Close()
}

type NatsWrapper struct {
	SC stan.Conn
}

func NewNatsWrapper() INatsWrapper {
	return &NatsWrapper{}
}

func (s *NatsWrapper) Connect(productionID string, clusterID string, URL string) error {
	sc, err := stan.Connect(productionID, clusterID, stan.NatsURL(URL))
	if err != nil {
		return err
	}
	s.SC = sc
	return nil
}

func (s *NatsWrapper) Publish(subject NatsEventType, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.SC.Publish(ParseNatsEventType(subject), bytes)
}

func (s *NatsWrapper) Log(name string, message string, level domain.Level) error {
	return s.Publish(LOGGING, domain.Log{Name: name, Date: time.Now(), Message: message, Level: level})
}

func (s *NatsWrapper) Subscribe(subject NatsEventType, queueName string, DurableKey string, callback func(data []byte) error) {
	s.SC.QueueSubscribe(ParseNatsEventType(subject), queueName, func(msg *stan.Msg) {
		bytes := msg.Data
		err := callback(bytes)
		if err != nil {
			return
		}
		msg.Ack()
	}, stan.DurableName(DurableKey), stan.SetManualAckMode())
}

func (s *NatsWrapper) Close() {
	s.SC.Close()
}
