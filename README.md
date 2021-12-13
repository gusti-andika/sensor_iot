# GO REST-API
Backend with rest API that will listen latest temperature information from mqtt broker.
Navigate to rest-api folder in terminal and run `go run main.go `

## PREREQUISITES
go version >= 1.15 

## PARAMS
```shell script
usage: go run main.go
          -host string
                server hostname or IP (default "broker.emqx.io")
          -port int
                server port (default 1883)
          -username string
                username (default "emqx")
          -password string
                password (default "public")
```

# CLIENT REACT JS
Client react js that will poll rest API in 2 seconds interval to get the latest temperature and display on html
To run open terminal under client directory and run `npm start`

# PYTHON SCRIPT
Script located at simulation/pub_sub_tcp.py that will simulate IOT devices and will randomnly publish temperature to mqtt

## PREREQUISITES
* It supports Python 3.4+

## INSTALLATION
```bash
pip install paho-mqtt
```

## Run
```bash
python pub_sub_tcp.py
``` 