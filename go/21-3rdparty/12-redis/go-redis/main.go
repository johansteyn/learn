package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go 3rd-party Library: go-redis")
	fmt.Println()

	// Can connect to Redis using options directly...
	//opts := &redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//}
	// Or create options by parsing a URL...
	url := "redis://localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		handleError("Error parsing URL", err)
	}
	client := redis.NewClient(opts)
	defer client.Close()
	fmt.Printf("Connected to redis: %+v\n", client)
	fmt.Println()

	fmt.Println("Setting values...")
	set1(client, "name", "Johan")
	set1(client, "age", 42)
	set1(client, "vaccinated", true)
	fmt.Println()

	fmt.Println("Setting values a different way...")
	set2(client, "name", "Johan")
	set2(client, "age", 42)
	set2(client, "vaccinated", true)
	fmt.Println()

	fmt.Println("Checking existence of values...")
	exists(client, "name")
	exists(client, "age")
	exists(client, "vaccinated")
	exists(client, "nonexistent")
	fmt.Println()

	fmt.Println("Getting values...")
	getString1(client, "name")
	getInt1(client, "age")
	getBool1(client, "vaccinated")
	fmt.Println()

	fmt.Println("Getting values a different way...")
	getString2(client, "name")
	getInt2(client, "age")
	getBool2(client, "vaccinated")
	fmt.Println()

	fmt.Println("Getting values a yet another way...")
	getString3(client, "name")
	getInt3(client, "age")
	getBool3(client, "vaccinated")
	fmt.Println()

	fmt.Println("Incrementing values...")
	fmt.Println("(TODO)")
	// Can only increment int values
	////incr(connection, "name")
	//incr(connection, "age")
	//getInt(connection, "age")
	//incrby(connection, "age", 5)
	//getInt(connection, "age")
	fmt.Println()

	fmt.Println("Decrementing values...")
	fmt.Println("(TODO)")
	// Can only decrement int values
	////decr(connection, "name")
	//decr(connection, "age")
	//getInt(connection, "age")
	//decrby(connection, "age", 7)
	//getInt(connection, "age")
	fmt.Println()

	fmt.Println("Expiring and persisting...")
	fmt.Println("(TODO)")
	//ttl(connection, "age")
	//expire(connection, "age", 2)
	//ttl(connection, "age")
	//pause(1)
	////pause(3)
	//persist(connection, "age")
	//ttl(connection, "age")
	//exists(connection, "age")
	fmt.Println()

	fmt.Println("Deleting value...")
	fmt.Println("(TODO)")
	//delete(connection, "vaccinated")
	//exists(connection, "vaccinated")
	fmt.Println()

	fmt.Println("Lists...")
	fmt.Println("(TODO)")
	//rpush(connection, "names", "Bob")
	//rpush(connection, "names", "Carol")
	//lpush(connection, "names", "Alice")
	//llen(connection, "names")
	//lrange(connection, "names", 0, -1)
	//lrange(connection, "names", 1, 2)
	//rpop(connection, "names")
	//lrange(connection, "names", 0, -1)
	//lpop(connection, "names")
	//lrange(connection, "names", 0, -1)
	fmt.Println()

	fmt.Println("Sets...")
	fmt.Println("(TODO)")
	//sadd(connection, "vegetarians", "Alice")
	//sadd(connection, "vegetarians", "Carol")
	//// Can add the same value multiple times - it's a set, so it's unique
	//sadd(connection, "vegetarians", "Alice")
	//sadd(connection, "clowns", "Bob")
	//sadd(connection, "clowns", "Carol")
	//sismember(connection, "vegetarians", "Alice")
	//sismember(connection, "vegetarians", "Bob")
	//sismember(connection, "vegetarians", "Carol")
	//sismember(connection, "vegetarians", "Dave")
	//sismember(connection, "clowns", "Alice")
	//sismember(connection, "clowns", "Bob")
	//sismember(connection, "clowns", "Carol")
	//sismember(connection, "clowns", "Dave")
	//smembers(connection, "vegetarians")
	//smembers(connection, "clowns")
	//smembers(connection, "nonexistent")
	//srem(connection, "clowns", "Carol")
	//smembers(connection, "clowns")
	//// Popping from a set removes a random member...
	//spop(connection, "vegetarians")
	//smembers(connection, "vegetarians")
	//// TODO: How can we do SUNION with redigo?
	fmt.Println()

	fmt.Println("Sorted sets...")
	fmt.Println("(TODO)")
	//zadd(connection, "ages", 42, "Alice")
	////zadd(connection, "ages", 65, "Bob")
	//zadd(connection, "ages", 21, "Carol")
	//zrange(connection, "ages", 0, -1)
	//zrange(connection, "ages", 1, 2)
	fmt.Println()

	fmt.Println("Hash maps...")
	fmt.Println("(TODO)")
	//hset(connection, "person:alice", "name", "Alice")
	//hset(connection, "person:alice", "age", 42)
	//hset(connection, "person:alice", "vegetarian", true)
	//hmset(connection, "person:bob", map[string]interface{}{
	//	"name":       "Bob",
	//	"age":        65,
	//	"vegetarian": false})
	//hget(connection, "person:alice", "name")
	//hget(connection, "person:alice", "age")
	//hget(connection, "person:alice", "vegetarian")
	//hgetall(connection, "person:alice")
	//hgetall(connection, "person:bob")
	//hincrby(connection, "person:alice", "age", 5)
	//hgetall(connection, "person:alice")
	//hdel(connection, "person:alice", "vegetarian")
	//hgetall(connection, "person:alice")
	fmt.Println()

	fmt.Println("Pipelining...")
	fmt.Println("(TODO)")
	//pipeline(connection)
	fmt.Println()

	fmt.Println("Done.")
}

// The set1 method uses the generic Do method, which is used to execute arbitrary/custom commands
func set1(client *redis.Client, key string, value interface{}) {
	fmt.Printf("Setting key '%s' to value: %v\n", key, value)
	_, err := client.Do(context.Background(), "set", key, value).Result()
	if err != nil {
		handleError(fmt.Sprintf("Error setting '%s'", key), err)
	}
}

// The set2 method uses the more specifc 'Set' method
func set2(client *redis.Client, key string, value interface{}) {
	fmt.Printf("Setting key '%s' to value: %v\n", key, value)
	_, err := client.Set(context.Background(), key, value, 0).Result()
	if err != nil {
		handleError(fmt.Sprintf("Error setting '%s'", key), err)
	}
}

func exists(client *redis.Client, key string) {
	cmd := client.Exists(context.Background(), key)
	exists, err := cmd.Result()
	if err != nil {
		handleError(fmt.Sprintf("Error checking if '%s' exists", key), err)
	}
	// For some reason Exists returns an IntCmd, not a BoolCmd...
	if exists == 1 {
		fmt.Printf("Key '%s' exists\n", key)
	} else {
		fmt.Printf("Key '%s' does not exist\n", key)
	}
}

// The get*1 methods use the generic Do method, which is used to execute arbitrary/custom commands,
// followed by calls to Text/Int/Bool methods
func getString1(client *redis.Client, key string) {
	value, err := client.Do(context.Background(), "GET", key).Text()
	if err != nil {
		handleError(fmt.Sprintf("Error getting value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has string value: %s\n", key, value)
}

func getInt1(client *redis.Client, key string) {
	value, err := client.Do(context.Background(), "GET", key).Int()
	if err != nil {
		handleError(fmt.Sprintf("Error getting value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has int value: %d\n", key, value)
}

func getBool1(client *redis.Client, key string) {
	value, err := client.Do(context.Background(), "GET", key).Bool()
	if err != nil {
		handleError(fmt.Sprintf("Error getting value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has bool value: %t\n", key, value)
}

// The get*2 methods use more specifc methods
func getString2(client *redis.Client, key string) {
	cmd := redis.NewStringCmd(context.Background(), "GET", key)
	client.Process(context.Background(), cmd)
	value, err := cmd.Result()
	if err != nil {
		handleError(fmt.Sprintf("Error getting string value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has string value: %s\n", key, value)
}

func getInt2(client *redis.Client, key string) {
	cmd := redis.NewIntCmd(context.Background(), "GET", key)
	client.Process(context.Background(), cmd)
	value, err := cmd.Result()
	if err != nil {
		handleError(fmt.Sprintf("Error getting int value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has int value: %d\n", key, value)
}

func getBool2(client *redis.Client, key string) {
	cmd := redis.NewBoolCmd(context.Background(), "GET", key)
	client.Process(context.Background(), cmd)
	value, err := cmd.Result()
	if err != nil {
		handleError(fmt.Sprintf("Error getting int value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has bool value: %t\n", key, value)
}

// The get*3 methods use the more specifc 'Get' method, which returns a string
func getString3(client *redis.Client, key string) {
	value, err := client.Get(context.Background(), key).Result()
	if err != nil {
		handleError(fmt.Sprintf("Error getting string value for key '%s'", key), err)
	}
	fmt.Printf("Key '%s' has string value: %s\n", key, value)
}

func getInt3(client *redis.Client, key string) {
	value, err := client.Get(context.Background(), key).Result()
	if err != nil {
		handleError(fmt.Sprintf("Error getting string value for key '%s'", key), err)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		handleError(fmt.Sprintf("Error converting string '%s' to int", value), err)
	}
	fmt.Printf("Key '%s' has int value: %d\n", key, intValue)
}

func getBool3(client *redis.Client, key string) {
	value, err := client.Get(context.Background(), key).Result()
	if err != nil {
		handleError(fmt.Sprintf("Error getting string value for key '%s'", key), err)
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		handleError(fmt.Sprintf("Error converting string '%s' to bool", value), err)
	}
	fmt.Printf("Key '%s' has bool value: %t\n", key, boolValue)
}

//func pause(seconds int) {
//	fmt.Printf("Sleeping for %d seconds...\n", seconds)
//	time.Sleep(time.Second * time.Duration(seconds))
//}

func handleError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
	os.Exit(1)
}
