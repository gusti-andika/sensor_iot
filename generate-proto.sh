#!/bin/bash

#generate python protobuf message
protoc ./proto/sensor.proto --python_out=./simulation

#generate golang protobuf message
protoc ./proto/sensor.proto --go_out=plugins=grpc:.