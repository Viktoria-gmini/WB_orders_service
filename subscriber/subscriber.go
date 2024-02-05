package subscriber

import (
	"fmt"
	"log"
	"sync"

	cacher "github.com/nats-io/go-nats-streaming/services/cacher"
	sender "github.com/nats-io/go-nats-streaming/services/sender"
	stan "github.com/nats-io/stan.go"
)

func Subscribe() {
	sc, err := stan.Connect("test-cluster", "test-client")
	if err != nil {
		log.Fatal(err)
	}
	cache := cacher.New(0, 0)

	defer sc.Close()
	sc.Subscribe("hello-subject", func(msg *stan.Msg) {
		fmt.Printf("Got: %s\n", string(msg.Data))
		err = sender.Confirm(msg.Data, cache)
		if err != nil {
			fmt.Println("2")
			panic(err)
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
