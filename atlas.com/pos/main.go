package main

import (
	"atlas-pos/kafka/consumers"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := log.New(os.Stdout, "pos ", log.LstdFlags|log.Lmicroseconds)

	createEventConsumers(l)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Println("[INFO] shutting down via signal:", sig)
}

func createEventConsumers(l *log.Logger) {
	cec := func(topicToken string, emptyEventCreator consumers.EmptyEventCreator, processor consumers.EventProcessor) {
		createEventConsumer(l, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_ENTER_PORTAL", consumers.PortalEnterCommandCreator(), consumers.HandlePortalEnterCommand())
}

func createEventConsumer(l *log.Logger, topicToken string, emptyEventCreator consumers.EmptyEventCreator, processor consumers.EventProcessor) {
	h := func(logger *log.Logger, event interface{}) {
		processor(logger, event)
	}

	c := consumers.NewConsumer(l, context.Background(), h,
		consumers.SetGroupId("Portal Service"),
		consumers.SetTopicToken(topicToken),
		consumers.SetEmptyEventCreator(emptyEventCreator))
	go c.Init()
}
