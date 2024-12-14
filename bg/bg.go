package bg

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/nikonor/alert-bot/event"
)

var (
	events = make([]event.Event, 0)
	l      sync.Mutex
)

func New(ctx context.Context) {
	go run(ctx)
}

func run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			check(time.Now())
		}
	}
}

func Add(e event.Event) {
	l.Lock()
	defer l.Unlock()

	events = append(events, e)

	sort.Slice(events, func(i, j int) bool {
		return events[i].T.Before(events[j].T)
	})
}

func check(now time.Time) {
	l.Lock()
	defer l.Unlock()

	var (
		idx int
		ev  event.Event
	)
	for idx, ev = range events {
		fmt.Printf("=>%s=>%s\n", ev.T.Format("15:04:05"), ev.Msg)
		if now.Unix() < ev.T.Unix() {
			break
		}
	}
	fmt.Printf("idx=%d,len=%d\n", idx, len(events))
	if idx > 0 {
		events = events[idx:]
	}
}
