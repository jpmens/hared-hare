`hared.go`: Golang version of reference python implementation.

* prerequisites:
```
   go get github.com/eclipse/paho.mqtt.golang
   go get gopkg.in/gcfg.v1
```
* run as: `go run hared.go`
* build a binary with: `go build hared.go`

The Go version of _hared_ also supports the `HARED_INI` configuration. If the
file cannot be opened/parsed, a diagnostic message is issued and the program
launches with its built-in defaults.

Defaults:
```
[defaults]
verbose    = True
udpHost    = 127.0.0.1
udpPort    = 8035
mqttURI    = tcp://127.0.0.1:1883
mqttTopic  = logging/hare
mqttQos    = 1
```

Complete `HARED_INI` example:
```
[defaults]
verbose    = False
udpHost    = localhost
udpPort    = 8035
# Where to publish to (port is optional):
# for plain MQTT:  tcp://127.0.0.1:1883
# for TLS MQTT:    ssl://hostname.example.com:8883
mqttURI    = tcp://127.0.0.1:1883
mqttTopic  = logging/hare
mqttQos    = 1
mqttUser   = username
mqttPass   = password
# path to CA certificate file if mqttURI starts with "tls://"
mqttCAfile = /etc/my/cert.pem
```
