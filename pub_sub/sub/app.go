package main

import (
	"context"
	"fmt"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"log"
	"net/http"
	"os"
)

var (
	sub = &common.Subscription{
		PubsubName: "pubsub",
		Topic:      "orders",
		Route:      "/orders",
	}
	appPort     string
	daprService common.Service
)

func init() {
	appPort = os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "6005"
	}

	// Create the new server on appPort and add a topic listener
	daprService = daprd.NewService(":" + appPort)
	err := daprService.AddTopicEventHandler(sub, eventHandler)
	if err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

}

func main() {
	// Start the server
	err := daprService.Start()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	fmt.Println("Subscriber received:", e.Data)
	return false, nil
}
