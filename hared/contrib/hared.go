package main

// Author: Giovanni Angoli <juzam76@gmail.com>

import (
    "os"
    "fmt"
    "net"
    "strconv"
    "io/ioutil"
    "crypto/tls"
    "crypto/x509"
    "encoding/json"
    "gopkg.in/gcfg.v1"
    "github.com/eclipse/paho.mqtt.golang"
)

func main() {

    cfg := struct {
        Defaults struct {
            Verbose    bool
            UdpHost    string
            UdpPort    int
            MqttURI    string
            MqttTopic  string
            MqttQos    byte
            MqttUser   string
            MqttPass   string
            MqttCAfile string
        }
    }{}

    cfg.Defaults.Verbose        = false
    cfg.Defaults.UdpHost        = "0.0.0.0"
    cfg.Defaults.UdpPort        = 8035
    cfg.Defaults.MqttURI        = "tcp://localhost:1883"
    cfg.Defaults.MqttTopic      = "logging/hare"
    cfg.Defaults.MqttQos        = 1

    cfgfile := "/usr/local/etc/hared.ini"

    if value, ok := os.LookupEnv("HARED_INI"); ok {
        cfgfile = value
    }

    error := gcfg.ReadFileInto(&cfg, cfgfile)
    if error != nil {
        fmt.Println(error)
    }

    if cfg.Defaults.Verbose {
        fmt.Println(cfg)
    }

    ServerAddr, _ := net.ResolveUDPAddr("udp", cfg.Defaults.UdpHost + ":" + strconv.Itoa(cfg.Defaults.UdpPort))
    ServerConn, _ := net.ListenUDP("udp", ServerAddr)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n,_,_ := ServerConn.ReadFromUDP(buf)

        message := string(buf[0:n])

        if cfg.Defaults.Verbose {
           fmt.Println(message)
        }

        var data map[string]interface{}
        if  error := json.Unmarshal([]byte(message), &data); error != nil {
            continue
        }

        opts := mqtt.NewClientOptions().AddBroker(cfg.Defaults.MqttURI)

        if len(cfg.Defaults.MqttUser) > 0 && len(cfg.Defaults.MqttPass) > 0 {
            opts.SetUsername(cfg.Defaults.MqttUser)
            opts.SetPassword(cfg.Defaults.MqttPass)
        }

        if len(cfg.Defaults.MqttCAfile) > 0 {
            CA_Pool := x509.NewCertPool()
            severCert, _ := ioutil.ReadFile(cfg.Defaults.MqttCAfile)
            CA_Pool.AppendCertsFromPEM(severCert)
            opts.SetTLSConfig(&tls.Config{RootCAs: CA_Pool})
        }

        c := mqtt.NewClient(opts)
        if token := c.Connect(); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
                    continue
        }
        if token := c.Publish(cfg.Defaults.MqttTopic, cfg.Defaults.MqttQos, false, message); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
        }
	c.Disconnect(250)
    }
}
