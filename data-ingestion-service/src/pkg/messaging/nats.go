package messaging

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

type NatsManager struct {
	nc *nats.Conn
}

func NewNatsManager(uri string) *NatsManager {
	nc, err := nats.Connect(uri)
	if err != nil {
		log.Fatal(err)
	}
	return &NatsManager{nc}
}

func (nm *NatsManager) GetClient() *nats.Conn {
	return nm.nc
}

func (nm *NatsManager) Close() {
	nm.nc.Close()
}

func (nm *NatsManager) Subscribe(channel string) *nats.Subscription {
	if nm.nc == nil {
		return nil
	}
	if nm.nc.IsClosed() {
		return nil
	}
	sub, err := nm.nc.SubscribeSync(channel)
	if err != nil {
		log.Error(err)
		return nil
	}
	return sub
}

func (nm *NatsManager) SubscribeWithCallback(channel string, callback func(*nats.Msg)) {
	if nm.nc == nil {
		return
	}
	if nm.nc.IsClosed() {
		return
	}
	_, err := nm.nc.Subscribe(channel, callback)
	if err != nil {
		log.Error(err)
		return
	}
}

func (nm *NatsManager) Publish(channel string, data []byte) error {
	return nm.nc.Publish(channel, data)
}
