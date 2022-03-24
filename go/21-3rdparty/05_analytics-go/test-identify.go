package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/analytics-go"
	"github.com/segmentio/ksuid"
)

func main() {
	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	id := ksuid.New()
	fmt.Printf("Generated  anonymous ID: %v\n", id)

	// Start by doing an anonymous track event call
	client.Enqueue(analytics.Track{
		Event: "test-identify",
		AnonymousId: id.String(),
	})

	// Pausing between Segment calls otherwise I've seen them in unexpected order on Segment...
	time.Sleep(1 * time.Second)

	// Then obtain a user's email address and do an identify call
	email := "johan.steyn@gmail.com"
	h := sha256.New()
	h.Write([]byte(email))
	auid := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("Derived AUID from email: %s\n", auid)
	traits := analytics.NewTraits()
	traits.SetEmail(email)
	traits.Set("SomeTrait", "SomeValue")
	traits.Set("AnotherTrait", "AnotherValue")
	traits.Set("YetAnotherTrait", "YetAnotherValue")
	client.Enqueue(analytics.Identify{
		// The anonymousId is not required
		// The JS and Ruby examples have only userId:
		//   https://segment.com/docs/getting-started/04-full-install
		AnonymousId: id.String(),
		// In fact, not even the userId is required... 
		// but you cannot omit both anonymousId and userId
		UserId: auid,
		Traits: traits,
	})

	time.Sleep(1 * time.Second)

	// Do another track event call, this time with both anonymous and user IDs
	client.Enqueue(analytics.Track{
		Event: "test-identify",
		AnonymousId: id.String(),
		UserId: auid,
	})

	client.Close()
}

