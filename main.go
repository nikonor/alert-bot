package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nikonor/alert-bot/bg"
	"github.com/nikonor/alert-bot/event"
)

func main() {
	ctx := context.Background()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	bg.New(ctx)

	for _, d := range []string{"1h", "10s", "1m", "30m", "3s"} {
		dur, _ := time.ParseDuration(d)
		fmt.Println("add event to " + dur.String())
		bg.Add(event.NewWDuration(dur, d))
	}

	<-sigChan
}
