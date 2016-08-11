package main

import (
	"fmt"
	"github.com/nats-io/nats"
	"sync/atomic"
	"time"
)

func main() {
	i := int32(0)

	nc, _ := nats.Connect(nats.DefaultURL)

	endChan := make(chan int)
	sub, _ := nc.SubscribeSync("Benchmark")
	sub.NextMsg(5 * time.Second)

	startTime := time.Now()

	nc.Subscribe("Benchmark", func(m *nats.Msg) {
		atomic.AddInt32(&i, 1)
		if i > 9000000 {
			endChan <- 1
		}
	})

	<-endChan
	dur := time.Since(startTime)
	fmt.Println(dur)
	fmt.Println(i)
}
