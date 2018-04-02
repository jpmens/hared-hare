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
            MqttClient string
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
    cfg.Defaults.MqttClient     = "hared"
    cfg.Defaults.MqttTopic      = "logging/hare"
    cfg.Defaults.MqttQos        = 1

    cfgfile := "/usr/local/etc/hared.ini"

    if value, ok := os.LookupEnv("HARED_INI"); ok {
        cfgfile = value
    }

    err := gcfg.ReadFileInto(&cfg, cfgfile)
    if err != nil {
        fmt.Println(err)
    }

    if cfg.Defaults.Verbose {
        fmt.Println(cfg)
    }

    ServerAddr, _ := net.ResolveUDPAddr("udp", cfg.Defaults.UdpHost + ":" + strconv.Itoa(cfg.Defaults.UdpPort))
    ServerConn, _ := net.ListenUDP("udp", ServerAddr)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n,address,_ := ServerConn.ReadFromUDP(buf)

        message := string(buf[0:n])
        remote := address.IP.String();

        if cfg.Defaults.Verbose {
           fmt.Println(message)
        }

        var data map[string]interface{}
        if  err := json.Unmarshal([]byte(message), &data); err != nil {
            continue
        }

        data["remote"] = remote
        data["tst"] =  int64(data["tst"].(float64))
        rawpayload, _ := json.Marshal(data)
        payload := string(rawpayload)

        if cfg.Defaults.Verbose {
           fmt.Println(payload)
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

        opts.SetClientID(cfg.Defaults.MqttClient)
        c := mqtt.NewClient(opts)
        if token := c.Connect(); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
                    continue
        }
        if token := c.Publish(cfg.Defaults.MqttTopic, cfg.Defaults.MqttQos, false, payload); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
        }
	c.Disconnect(250)
    }
}
