module github.com/gusti-andika/sensor_iot/rest-api

go 1.15

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	google.golang.org/protobuf v1.27.1
)

replace github.com/gusti-andika/sensor_iot/mymqtt => ./mymqtt
