# python 3.6

import json
import random
import time
import proto.sensor_pb2

from paho.mqtt import client as mqtt_client

from proto.sensor_pb2 import Temperature


BROKER = 'w7de211b.us-east-1.emqx.cloud'
PORT = 15301
TOPIC = "temperature/latest"
# generate client ID with pub prefix randomly
CLIENT_ID = "python-mqtt-tcp-pub-sub-{id}".format(id=random.randint(0, 1000))
USERNAME = 'emqx'
PASSWORD = 'public'
FLAG_CONNECTED = 0


def on_connect(client, userdata, flags, rc):
    global FLAG_CONNECTED
    if rc == 0:
        FLAG_CONNECTED = 1
        print("Connected to MQTT Broker!")
        client.subscribe(TOPIC)
    else:
        print("Failed to connect, return code {rc}".format(rc=rc), )


def on_message(client, userdata, msg):
    print("Received `{payload}` from `{topic}` topic".format(
        payload=msg.payload.decode(), topic=msg.topic))


def connect_mqtt():
    client = mqtt_client.Client(CLIENT_ID)
    client.username_pw_set(USERNAME, PASSWORD)
    client.on_connect = on_connect
    #client.on_message = on_message
    client.connect(BROKER, PORT)
    return client


def publish(client):
    msg_count = 0
    while True:
        temp = Temperature()
        temp.value = random.randint(25, 35)

        result = client.publish(TOPIC, temp.SerializeToString())
        # result: [0, 1]
        status = result[0]
        if status == 0:
            print("Send `{msg}` to topic `{topic}`".format(msg=temp, topic=TOPIC))
        else:
            print("Failed to send message to topic {topic}".format(topic=TOPIC))
        msg_count += 1
        time.sleep(2)


def run():
    client = connect_mqtt()
    client.loop_start()
    time.sleep(1)
    if FLAG_CONNECTED:
        publish(client)
    else:
        client.loop_stop()


if __name__ == '__main__':
    run()
