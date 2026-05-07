package constants

import "time"

const (
	ShutdownDelay         = 5 * time.Second
	ServerShutdownTimeout = 10 * time.Second
	HardShutdownWait      = 5 * time.Second
)
