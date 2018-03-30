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

Extra `HARED_INI` fields supported:
- mqttuser
- mqttpass
