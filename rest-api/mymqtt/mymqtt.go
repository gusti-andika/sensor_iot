package mymqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MyMqtt struct {
	client mqtt.Client
}

type MessageHandler func(message []byte)

func (m *MyMqtt) Connect(config Config) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Host, config.Port))
	opts.SetUsername(config.Username) // these are the credentials that you declare for your cluster
	opts.SetPassword(config.Password)

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	// opts.DefaultPublishHandler = messagePubHandler
	// create the client using the options above
	client := mqtt.NewClient(opts)

	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	m.client = client
}

func (m *MyMqtt) Subscribe(topic string, handler MessageHandler) {
	// subscribe to the same topic, that was published to, to receive the messages
	token := m.client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		handler(msg.Payload())
	})
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic\n")
		panic(token.Error())
	}

	fmt.Printf("Subscribed to topic: %s\n", topic)
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}
