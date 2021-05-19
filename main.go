package main

import (
	"log"
	"temporal/activities"
	"temporal/worflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "GREETING_TASK_QUEUE", worker.Options{})
	w.RegisterWorkflow(worflows.Greetings)
	w.RegisterActivity(activities.Greet)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
