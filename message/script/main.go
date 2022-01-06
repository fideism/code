package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fideism/code/message"
)

func main() {
	entity, err := message.NewMqtt(message.MqttConfig{
		Client:                "client_name",
		UserName:              "username",
		Password:              "password",
		Broker:                "192.168.2.210:1883",
		ReconnectingHandler:   message.DefaultReconnectingHandler,
		ConnectionLostHandler: message.DefaultConnectionLostHandler,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(entity)

	if token := entity.Subscribe(`message`, 0x01, message.DefaultSubscribeHandler); token.Wait() && token.Error() != nil {

		fmt.Println(token.Error())
		os.Exit(1)
	}

	//topics := make(map[string]byte, 10)
	//for i := 0; i < 10; i++ {
	//	mac := fmt.Sprintf(`topic_%d`, i)
	//	topics[mac] = 0x01
	//}
	//entity.KeepSubscribeMultiple(topics, message.DefaultSubscribeHandler)

	for i := 0; i < 10; i++ {
		if token := entity.Publish(`message`, 0x01, true, fmt.Sprintf("message_%d", i)); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		time.Sleep(time.Second)
	}

	//entity.KeepPublish(`topic_5`, 0x01, true, "hello")
}
