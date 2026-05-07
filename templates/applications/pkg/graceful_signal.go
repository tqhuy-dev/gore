package pkg

import (
	"context"
	"os/signal"
	"syscall"
)

func GetSignalCtx() (signalCtx context.Context, signalCtxStop context.CancelFunc) {
	return signal.NotifyContext(context.Background(),
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGQUIT, // Ctrl+\
		syscall.SIGTERM, // the normal way to politely ask a program to terminate
	)
}
