https://en.wikipedia.org/wiki/NATS_Messaging
https://nats.io

https://github.com/nats-io/nats.go
https://pkg.go.dev/github.com/nats-io/nats.go

https://natsbyexample.com

https://www.youtube.com/watch?v=hjXIUPZ7ArM

MATS messaging is split in 2 parts:
 1. Core: Standard best-effort, at-most-once messaging
 2. Jetstream: Layered on top of Core to add persistence (key-value or object store) required for streaming and adding at-least-once and exactly-once messaging.
Easiest way to get started is with the CLI tool: has built-in servers and all the features to simulate a NATS environment
There are over 40 client libraries for NATS, and just about everything they can do can be done with the CLI tool.
Various ways to install NATS - he uses a Mac so uses homebrew:
  $ brew install nats-io/nats-tools/nats
Then, to start the NATS server:
  $ nats server run
That creates some accounts you can use to access the server.
To use the just-created "nats_development" user for future calls:
  $ nats context select nats_development
To do a simple request-reply pattern:
  $ nats reply hello.johan "Hello Johan :)"
NATS uses the concept of subject-based addressing: you can publish/subscribe to a series of tokens separated by periods.
Now we can make a request to the hello.johan subject:
  $ nats req hello.johan ""
In the above request the payload string is empty and simply ignored, but it can be anything: simple text, JSON, BSON, etc.
NATS is payload agnostic, which makes it very generic.
The requestor and replier are decoupled, ie. they don't need to know about each other - they only need to know the subject they want to talk about.
Synadia believes all communication can be boiled down to 2 constructs:
 1. Services, like the request-response example above
 2. Streams, allowing temporal decoupling between publishers and subscribers
For streaming, we subscribe to a subject:
  $ nats sub hello.world
Now we can publish to the hello.world subject:
  $ nats pub hello.world "Johan"
We can do a "fan-in" pattern with one subscriber and multiple publishers:
  $ nats pub hello.world "Johan" --count=10 --sleep=1s
To publish indefinitely, use a cout value of -1:
  $ nats pub hello.world "Johan" --count=-1 --sleep=1s
To do a "fan-out" pattern, create multiple subscribers:
  $ nats sub hello.world
  $ nats sub hello.world
  $ nats sub hello.world
With the "fan-out" pattern you can use "queue groups", placing multiple subscribers on a subject in a group together.
NATS will evenly distribute the load between the subscribers, which is handy for horizontal scalability - no need to set up your own, separate load balancer :)
  $ nats sub hello.world  --queue greeter <= In one terminal
  $ nats sub hello.world  --queue greeter <= In another terminal
Now send a thousand messages without sleeping:
  $ nats pub hello.world "Johan" --count=1000
In my case the first subscriber received 460 messages while the second one received 560.
Sending a million messages, the first subscriber received 500391 while the second one received 499609.
Subjects can be specified using wildcards:
  $ nats sub hello.*        <= matches subjects with first token "hello" and any second token (but cannot have three or more tokens)
  $ nats sub first.*.third  <= matches subjects with first token "first", any second token, and third token "third" (but cannot have 4 or more tokens)
  $ nats sub first.>        <= matches subjects with first token "first" and any subsequent tokens.
  $ nats sub >              <= matches ALL subjects (called "wiretap mode")


https://docs.nats.io/nats-concepts/jetstream

The JetStream persistence engine allows messages to be stored and replayed later, thereby enabling temporal decoupling.
Temporal decoupling can also be achieved with durable subscribers and queues, but neither is perfect:
  - Durable subscribers need to be created before messages are published.
  - Queue are meant for workload distribution and consumption - not for message replay
Streaming requires persistency so that messages can be replayed on demand as many times as required.
Replay policies:
  - Complete replay of ALL messages currently persisted, either instantly (as fast as they can be consumed) or at the original rate that they were published
  - Messages starting from a specified sequence number
  - Messages start from a specified start time
  - Only the last persisted message, or the last message for each subject.
Retention policies:
  Messages need not persist forever, ie. streams can't always grow forever - there needs to be limits on the sizes of streams.
  Types of limits:
  - Maximum message age 
  - Maximum number of messages
  - Maximum individual message size
  - Maximum stream size (in bytes)
  - Number of consumers
Discard policies:
  - Discard old <= Delete the oldest messages to make room for new messages
  - Discard new <= Reject any new messages, returning an error indicating that a limit was reached, until there is room for new messages
Retention policies:
  - Limit       <= Provide a replay of messages in the stream (the default retention policy)
  - Work queue  <= Provide exactly-once consumption where messages are removed from the shared queue as they are consumed
  - Interest    <= Variation of work queue where messages are only kept as long as there are consumers defined on the stream that haven't consumed the messages yet.
JetStream can apply transformations to messages as they are ingested into a stream (ie. mapping)
Persistent storage types:
  - Memory
  - File
  - Replication (for fault tolerance)
JetStream uses a RAFT distributed quorum algorithm to distribute persistency, while maintaining immediate consistency even in the face of failures (as opposed to eventual consistency)



