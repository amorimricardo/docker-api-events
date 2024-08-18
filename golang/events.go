package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	// Set the client to use API version 1.45
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.45"))
	if err != nil {
		log.Fatal(err)
	}

	// Create a filter to only get "die" events
	eventFilter := filters.NewArgs()
	eventFilter.Add("event", "die")

	// Get the docker events with the filter
	msgs, errs := cli.Events(context.Background(), events.ListOptions{Filters: eventFilter})

	// Loop through the messages and errors
	for {
		select {
		case msg := <-msgs:
			// Print container name, ID, and time for "die" events
			if msg.Action == "die" {
				containerName := msg.Actor.Attributes["name"]
				containerID := msg.ID
				eventTime := time.Unix(msg.Time, 0).Format(time.RFC3339)
				fmt.Printf("Container %s with ID %s died at %s\n", containerName, containerID, eventTime)
			}
		case err := <-errs:
			log.Fatal(err)
		}
	}
}
