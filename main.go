package main

import (
	"sync"

	cacher "github.com/nats-io/go-nats-streaming/services/cacher"
	subscriber "github.com/nats-io/go-nats-streaming/sub-pub/subscriber"
	web_service "github.com/nats-io/go-nats-streaming/web/web-service"
)

func main() {
	cache := cacher.New(0, 0)
	go subscriber.Subscribe(cache)
	go web_service.WebService(cache)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
