package main

import (
	"context"
	"fmt"
	"os"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func main() {
	// Connect
	topicName := "testeventtopic"
	subscriptionName := "s1"
	connStr := "Endpoint=sb://choreodev1.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=JFB12zRbyy4zeQ11pYXe1o+/qdGTzEnrRzPDhiKpmwc="
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
	fmt.Println("Starting 1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t, err := getTopic(ns, topicName)
	if err != nil {
		fmt.Printf("failed to build a new topic named %q\n", topicName)
		os.Exit(1)
	}

	subManager := t.NewSubscriptionManager()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	entity, err := subManager.Get(ctx, subscriptionName)

	fmt.Println(entity.Entity.Name)

	// subscription, err := t.NewSubscription(subscriptionName, func(sub *servicebus.Subscription) error {
	// 	return errors.New("math: square root of negative number")
	// })
	subscription, err := t.NewSubscription(subscriptionName)
	if subscription == nil {
		fmt.Println("NewSubscription is null")
		os.Exit(1)
	}
	fmt.Println("after NewSubscription")
	fmt.Println(subscription.Name)

	err = subscription.ReceiveOne(ctx, servicebus.HandlerFunc(func(ctx context.Context, message *servicebus.Message) error {
		fmt.Println(string(message.Data))
		return message.Complete(ctx)
	}))

	if err != nil {
		fmt.Println("FATAL: ", err)
		return
	}

	// type Topic struct {
	// 	// contains filtered or unexported fields
	// }

	// type Subscription struct {
	// 	Topic *Topic
	// 	// contains filtered or unexported fields
	// }

	//subs, err := t.NewSubscription(subscriptionName, SubscriptionWithReceiveAndDelete())

	// listenHandle, err := subs.Receive(ctx, func(ctx context.Context, message *servicebus.Message) servicebus.DispositionAction {
	// 	text := string(message.Data)
	// 	if text == "exit\n" {
	// 		fmt.Println("Oh snap!! Someone told me to exit!")
	// 		exit <- *new(struct{})
	// 	} else {
	// 		fmt.Println(string(message.Data))
	// 	}
	// 	return message.Complete()
	// })

	// handler := func(ctx context.Context, message *servicebus.Message) servicebus.DispositionAction {
	// 	text := string(message.Data)
	// 	if text == "exit\n" {
	// 		fmt.Println("Oh snap!! Someone told me to exit!")
	// 		exit <- *new(struct{})
	// 	} else {
	// 		fmt.Println(string(message.Data))
	// 	}
	// 	return message.Complete()
	// }

	// listenHandle, err := subs.Receive(ctx, handler)

	// defer listenHandle.Close(context.Background())

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println("I am listening...")

	// select {
	// case <-exit:
	// 	fmt.Println("closing after 2 seconds")
	// 	select {
	// 	case <-time.After(2 * time.Second):
	// 		listenHandle.Close(context.Background())
	// 		return
	// 	}
	// }

}

func getTopic(ns *servicebus.Namespace, topicName string) (*servicebus.Topic, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	subManagerStruct, err := ns.NewSubscriptionManager(topicName)

	return subManagerStruct.Topic, err

	// tm := ns.NewTopicManager()

	// te, err := tm.Get(ctx, topicName)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("after get ")
	// fmt.Println(te.Entity.Name)
	// if te == nil {
	// 	_, err := tm.Put(ctx, topicName)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// t, err := ns.NewTopic(topicName, func(t *servicebus.Topic) error {
	// 	return errors.New("math: square root of negative number")
	// })
	// return t, err
}
