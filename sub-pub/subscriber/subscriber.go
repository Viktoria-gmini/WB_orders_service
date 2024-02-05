package subscriber

import (
	"fmt"
	"log"
	"sync"

	cacher "github.com/nats-io/go-nats-streaming/services/cacher"
	sender "github.com/nats-io/go-nats-streaming/services/sender"
	stan "github.com/nats-io/stan.go"
)

var cache *cacher.Cache

func Subscribe(myCache *cacher.Cache) {
	sc, err := stan.Connect("test-cluster", "test-client")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	cache = myCache
	cache.UploadFromDB()

	sc.Subscribe("hello-subject", func(msg *stan.Msg) {
		fmt.Printf("Got: %s\n", string(msg.Data))
		err = sender.Confirm(msg.Data, cache)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Already in database! Congratulations")
	})

	Block()

}

func Block() {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
