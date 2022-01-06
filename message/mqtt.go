package message

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MqttConfig mqtt config
type MqttConfig struct {
	Client                string
	UserName              string
	Password              string
	Broker                string
	ReconnectingHandler   func(c mqtt.Client, options *mqtt.ClientOptions)
	ConnectionLostHandler func(client mqtt.Client, err error)
}

// MqttEntity entity
type MqttEntity struct {
	Client mqtt.Client
}

// NewMqtt new mqtt
func NewMqtt(cfg MqttConfig) (*MqttEntity, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetConnectTimeout(time.Second * 3)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetMaxReconnectInterval(5 * time.Second)
	opts.SetClientID(cfg.Client)
	opts.SetUsername(cfg.UserName)
	opts.SetPassword(cfg.Password)
	opts.SetKeepAlive(time.Minute)
	opts.SetConnectionLostHandler(cfg.ConnectionLostHandler)
	opts.SetReconnectingHandler(cfg.ReconnectingHandler)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MqttEntity{Client: client}, nil
}

// DefaultReconnectingHandler default func
func DefaultReconnectingHandler(c mqtt.Client, options *mqtt.ClientOptions) {
	fmt.Println("default reconnection handler")
}

// DefaultConnectionLostHandler default func
func DefaultConnectionLostHandler(client mqtt.Client, err error) {
	fmt.Printf("default connection lost handler, err:%s\n", err.Error())
}

// DefaultSubscribeHandler default func
func DefaultSubscribeHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Println("subscribe func")
	fmt.Println("message topic:", message.Topic())
	fmt.Println("message id:", message.MessageID())
	fmt.Println("message:", string(message.Payload()))
	fmt.Printf("default subscribe handler, message:%s\n", string(message.Payload()))
	message.Ack()
}

// Publish publish
func (e *MqttEntity) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return e.Client.Publish(topic, qos, retained, payload)
}

// KeepPublish keep publish
func (e *MqttEntity) KeepPublish(topic string, qos byte, retained bool, payload interface{}) error {
	token := e.Client.Publish(topic, qos, retained, payload)

	token.Wait()

	return token.Error()
}

// Subscribe subscribe
func (e *MqttEntity) Subscribe(topic string, qos byte, callback func(client mqtt.Client, message mqtt.Message)) mqtt.Token {
	return e.Client.Subscribe(topic, qos, callback)
}

// SubscribeMultiple subscribe multiple
func (e *MqttEntity) SubscribeMultiple(filters map[string]byte, callback func(client mqtt.Client, message mqtt.Message)) mqtt.Token {
	return e.Client.SubscribeMultiple(filters, callback)
}

// KeepSubscribe keep subscribe
func (e *MqttEntity) KeepSubscribe(topic string, qos byte, callback func(client mqtt.Client, message mqtt.Message)) {
	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGTERM)

	token := e.Client.Subscribe(topic, qos, callback)
	token.Wait()

	<-keepAlive
}

// KeepSubscribeMultiple keep subscribe multiple
func (e *MqttEntity) KeepSubscribeMultiple(filters map[string]byte, callback func(client mqtt.Client, message mqtt.Message)) {
	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGTERM)

	token := e.Client.SubscribeMultiple(filters, callback)
	token.Wait()

	<-keepAlive
}
