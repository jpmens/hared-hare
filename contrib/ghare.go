/*
 * hare.go (C)2018 by Jan-Piet Mens <jp@mens.de>
 * inspired by work done on hared.go by Juzam
 * This is a rewrite of hare.c in Golang
 * This is also the first time I've spent more than 10m with the language.
 */

package main

import (
       "encoding/json"
       "strings"
       "fmt"
       "time"
       "os"
       "net"
    )

const (
      PORTNUM = 8035
)

func main() {
	m := []string{"user", "service", "rhost", "tty"}
	var val string
	var ok bool

	val, ok = os.LookupEnv("PAM_SM_TYPE")	// FreeBSD
	if !ok {
	   val, ok = os.LookupEnv("PAM_TYPE")
	   if !ok {
	      // crap
	   }
	}

	if val != "open_session" && val != "pam_sm_open_session" {
	    fmt.Printf("%s: Neither PAM open_session nor pam_sm_open_session detected\n", os.Args[0])
	    return
	}

	hostname, err := os.Hostname()
	if err != nil {
	     fmt.Printf("Cannot get local hostname: %v\n", err)
	     return
	}

	address := "127.0.0.1"
	if len(os.Args) == 2 {
	   address = os.Args[1]
	}

	jlist  := make(map[string]interface{})

	for i := range m {
		k := m[i]
		p := "PAM_" + strings.ToUpper(k)
		v, ok := os.LookupEnv(p)
		if !ok {
		    jlist[k] = nil
		} else {
		  jlist[k] = v
		}
	}

	jlist["hostname"] = hostname
	jlist["tst"] = time.Now().Unix()

	jbin, err := json.Marshal(jlist)
	if err != nil {
	   fmt.Println("ERR")
	}
	fmt.Println(string(jbin))

	a := fmt.Sprintf("%s:%d", address, PORTNUM)

	conn, err := net.Dial("udp", a)
	if err != nil {
	   fmt.Printf("Cannot create UDP socket %v\n", err)
	   return
	}

	fmt.Fprintf(conn, "%s", string(jbin))

	conn.Close()
}