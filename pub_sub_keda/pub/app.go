package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	pubsubComponentName = "pubsub"
	pubsubTopic         = "orders"
)

var (
	publisherClient dapr.Client
)

func init() {
	// Create a new client for Dapr using the SDK
	var err error
	publisherClient, err = dapr.NewClient()
	if err != nil {
		panic(err)
	}

}

func main() {

	c := context.Background()

	// Publish events using Dapr pubsub
	for i := 1; i <= 10; i++ {
		order := `{"orderId":` + strconv.Itoa(i) + `}`

		err := publisherClient.PublishEvent(c, pubsubComponentName, pubsubTopic, []byte(order))
		if err != nil {
			panic(err)
		}
		fmt.Println("Published data:", order)

		time.Sleep(time.Second)
	}

	// https://docs.dapr.io/operations/hosting/kubernetes/kubernetes-job/
	// When running a basic Kubernetes Job, you need to call the /shutdown endpoint for the sidecar to gracefully stop and the job to be considered Completed.
	defer publisherClient.Close()
	defer publisherClient.Shutdown(c)
}
