package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const subject = "my.test.subject"

func main() {
	conn, _ := nats.Connect("nats://127.0.0.1:4222")
	defer conn.Close()

	for i := 1; ; i++ {
		msg := fmt.Sprintf("Zdarzenie numer %d", i)
		conn.Publish(subject, []byte(msg))
		log.Printf("Wyemitowano zdarzenie numer %d", i)
		time.Sleep(time.Second * 1)
	}
}
