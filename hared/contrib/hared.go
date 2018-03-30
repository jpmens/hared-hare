package main

// Author: Giovanni Angoli <juzam76@gmail.com>

import (
    "os"
    "fmt"
    "net"
    "strconv"
    "encoding/json"
    "gopkg.in/gcfg.v1"
    "github.com/eclipse/paho.mqtt.golang"
)

func main() {

    cfg := struct {
        Defaults struct {
            Verbose    bool
            Listenhost string
            Listenport int
            Mqtthost   string
            Mqttport   int
            Topic      string
            Mqttuser   string
            Mqttpass   string
        }
    }{}

    cfg.Defaults.Verbose    = false
    cfg.Defaults.Listenhost = "0.0.0.0"
    cfg.Defaults.Listenport = 8035
    cfg.Defaults.Mqtthost   = "localhost"
    cfg.Defaults.Mqttport   = 1883
    cfg.Defaults.Topic      = "logging/hare"

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

    ServerAddr, _ := net.ResolveUDPAddr("udp", cfg.Defaults.Listenhost + ":" + strconv.Itoa(cfg.Defaults.Listenport))
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

        opts := mqtt.NewClientOptions().AddBroker("tcp://" + cfg.Defaults.Mqtthost + ":" + strconv.Itoa(cfg.Defaults.Mqttport))

        if len(cfg.Defaults.Mqttuser) > 0 && len(cfg.Defaults.Mqttpass) > 0 {
            opts.SetUsername(cfg.Defaults.Mqttuser)
            opts.SetPassword(cfg.Defaults.Mqttpass)
        }

        c := mqtt.NewClient(opts)
        if token := c.Connect(); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
        }
        if token := c.Publish(cfg.Defaults.Topic, 0, false, message); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
        }
    }
}
