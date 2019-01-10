package main

import (
	"github.com/nats-io/go-nats-streaming"
	"fmt"
	"github.com/nats-io/go-nats-streaming/pb"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	sc, err := stan.Connect("test-cluster", "testClient2", stan.NatsURL("127.0.0.1:4222"))
	if err != nil{
		log.Fatalf(err.Error())
	}

	go func() {
		for i:=0; i< 1; i++ {
			fmt.Println("sending: ", "Hello_", string(i))
			sc.Publish("events", []byte("Hello_" + string(i)));
			time.Sleep(1 * time.Second)
		}
	}()
	mcb := func(msg *stan.Msg) {
		fmt.Println(msg)
	}

	startOpt := stan.StartAt(pb.StartPosition_NewOnly)
	sub, err := sc.QueueSubscribe("events", "1", mcb, startOpt, stan.DurableName("consumer2"))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}


	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	fmt.Println("waiting for interrupt");
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sub.Close()
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
