package utils

import (
	"os"
	"context"
	"os/signal"
)

func ContextWithCancelSignals(sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, sig...)
	go func() {
		<-exit
		cancel()
	}()
	return ctx
}
