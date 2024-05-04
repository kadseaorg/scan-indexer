package indexer

import (
	"context"
	"time"
)

func forever(ctx context.Context, interval time.Duration, runner func()) {
	if interval < time.Second {
		interval = time.Second
	}
	runner()
	timer := time.NewTimer(interval)
	defer func() {
		timer.Stop()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			runner()
			timer.Reset(interval)
		}
	}
}
