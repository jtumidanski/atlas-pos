package main

import (
	"atlas-pos/kafka/consumers"
	"atlas-pos/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := logger.CreateLogger()

	consumers.CreateEventConsumers(l)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Println("[INFO] shutting down via signal:", sig)
}
