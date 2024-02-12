package main

import (
	"fmt"

	"github.com/itzmeanjan/pubsub"
)

func main() {
	fmt.Println("Go 3rd-party Library: itzmeanjan/pubsub")
	fmt.Println()

	fmt.Println("Creating a new broker with one shard...")
	broker := pubsub.New(12)
	fmt.Printf("--- broker: %+v\n", broker)
	// TODO: Multiple shards...

	fmt.Println("Subscribing to two topics...")
	subscriber := broker.Subscribe(16, "topic_1", "topic_2")
	if subscriber == nil {
		fmt.Printf("❌ Failed to subscribe to topics\n")
		return
	}

	fmt.Println("Unsubscribing from first topic...")
	if subscriber.Unsubscribe("topic_1") == 1 {
		fmt.Printf("✅ Unsubscribed from `topic_1`\n")
	}

	fmt.Println("Subscribing to third topic...")
	if subscriber.AddSubscription("topic_3") == 1 {
		fmt.Printf("✅ Subscribed to `topic_3`\n")
	}

	fmt.Println("Publishing...")
	m := "hello"
	message := pubsub.Message{
		Topics: []string{
			"topic_1",
			"topic_2",
			"topic_3",
		},
		Data: []byte(m),
	}
	//fmt.Printf("✅ Published '%s' message to %d topics\n", m, broker.Publish(&message))
	count := broker.Publish(&message)
	fmt.Printf("✅ Published '%s' message to %d topics\n", m, count)
	fmt.Printf("--- broker: %+v\n", broker)

	fmt.Println("Listening for messages...")
	// Consume message by calling the "Next" method in a loop
	for range subscriber.Listener() {
		msg := subscriber.Next()
		if msg == nil {
			fmt.Println("No more messages - breaking out...")
			break
		}
		fmt.Printf("✅ Received `%s` message on topic `%s`\n", msg.Data, msg.Topic)
		if !subscriber.Consumable() {
			fmt.Println("Subscriber not consumable - breaking out...")
			break
		}
	}
	fmt.Println("Checking for any buffered messages...")
	if subscriber.Consumable() {
		fmt.Printf("❌ Somehow there are still some buffered messages...\n")
	} else {
		fmt.Printf("✅ Consumed all messages\n")
	}

	fmt.Println("Unsubscribing from all topics...")
	count = subscriber.UnsubscribeAll()
	fmt.Printf("✅ Unsubscribed from %d topics\n", count)

	fmt.Println("Destroying subscriber...")
	subscriber.Destroy()

	fmt.Println()
}
