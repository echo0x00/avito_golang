package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	ctx, cancelFunc := context.WithCancel(context.Background())
	signal.NotifyContext(ctx, os.Interrupt)

	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout, cancelFunc)
	if err := client.Connect(); err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer client.Close()

	fmt.Fprintf(os.Stderr, "...Connected to %s\n", address)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := client.Receive()
		if err != nil {
			log.Fatalf("Receive error: %v", err)
		}
		fmt.Fprintf(os.Stderr, "...Connection was closed by peer\n")
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		err := client.Send()
		if err != nil {
			log.Fatalf("Send error: %v", err)
		}
		fmt.Fprintf(os.Stderr, "...EOF\n")
	}()

	wg.Wait()

	select {
	case <-ctx.Done():
		log.Println("done by context")
		cancelFunc()
		client.Close()
	default:
	}
}
