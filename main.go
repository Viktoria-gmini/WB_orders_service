package main

import (
	"sync"

	cacher "github.com/nats-io/go-nats-streaming/services/cacher"
	subscriber "github.com/nats-io/go-nats-streaming/sub-pub/subscriber"
	web_service "github.com/nats-io/go-nats-streaming/web/web-service"
)

func main() {
	//инициализируем кэш
	cache := cacher.New(0, 0)
	//подписываемся на канал
	go subscriber.Subscribe(cache)
	//запускаем веб сервис
	go web_service.WebService(cache)
	//удерживаем main поток
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
