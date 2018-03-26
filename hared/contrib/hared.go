package main

// Author: Giovanni Angoli <juzam76@gmail.com>

import (
    "fmt"
    "net"
    "encoding/json"
    "github.com/eclipse/paho.mqtt.golang"
)

type Data struct {
    Host string `json:"host"`
    Msg string  `json:"msg"`
}

func main() {
    ServerAddr, _ := net.ResolveUDPAddr("udp",":8035")
    ServerConn, _ := net.ListenUDP("udp", ServerAddr)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n,addr,_ := ServerConn.ReadFromUDP(buf)

        message := string(buf[0:n]);
        address := addr.IP.String();

        data := Data{Host: address, Msg: message}

        js, _ := json.Marshal(data)
        fmt.Println(string(js))

        opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.130:1883")
        c := mqtt.NewClient(opts)
        if token := c.Connect(); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
        }
        if token := c.Publish("logging/hare", 0, false, js); token.Wait() && token.Error() != nil {
                    fmt.Println(token.Error())
        }
    }
}
