package main

import (
	// "fmt"
	"log"

	turboGenerator "github.com/nats-io/go-nats-streaming/services/generator"
	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

func main() {
	// подключение к локальному NATS Server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// подключение к NATS Streaming
	sc, err := stan.Connect("test-cluster", "test-client2", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	var jsonFile []byte
	// генератор Json
	jsonFile, err = turboGenerator.GenerateJSON()
	if err != nil {
		panic(err)
	}
	// публикация сообщения
	err = sc.Publish("hello-subject", jsonFile)
	if err != nil {
		log.Fatal(err)
	}

}
