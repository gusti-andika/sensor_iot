package mymqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MyMqtt struct {
	client mqtt.Client
}

type MessageHandler func(message []byte)

func (m *MyMqtt) Init() {
	var broker = "069c17e1d82d482b96938095847c6b0d.s1.eu.hivemq.cloud" // find the host name in the Overview of your cluster (see readme)
	var port = 8883                                                    // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("iot")     // set a name as you desire
	opts.SetUsername("gandika") // these are the credentials that you declare for your cluster
	opts.SetPassword("Abcd1234")

	// (optionally) configure callback handlers that get called on certain events
	//messageHandler := func(client mqtt.Client, msg mqtt.Message) {
	//	handler(msg.Payload())
	//}

	//opts.SetDefaultPublishHandler(messageHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// create the client using the options above
	client := mqtt.NewClient(opts)

	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	m.client = client

	//subscribe(topic, client)
}

func (m *MyMqtt) Subscribe(topic string, handler MessageHandler) {
	// subscribe to the same topic, that was published to, to receive the messages
	handlerWrapper := func(client mqtt.Client, msg mqtt.Message) {
		handler(msg.Payload())
	}

	token := m.client.Subscribe(topic, 1, handlerWrapper)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic\n")
		panic(token.Error())
	}

	fmt.Printf("Subscribed to topic: %s\n", topic)
}

// this callback triggers when a message is received, it then prints the message (in the payload) and topic
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// upon connection to the client, this is called
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func publish(client mqtt.Client) {
	// publish the message "Message" to the topic "topic/test" 10 times in a for loop
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
		if token.Error() != nil {
			fmt.Printf("Failed to publish to topic")
			panic(token.Error())
		}
		time.Sleep(time.Second)
	}
}
