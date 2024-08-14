package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	redigo "github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("Go 3rd-party Library: redigo")
	fmt.Println()

	address := ":6379"
	connection, err := redigo.Dial("tcp", address, redigo.DialPassword(""))
	if err != nil {
		handleError("Error connecting to Redis", err)
	}
	defer connection.Close()
	fmt.Printf("Connected to redis: %+v\n", connection)
	fmt.Println()

	fmt.Println("Cleaning up...")
	cleanup(connection)
	fmt.Println()

	fmt.Println("Setting values...")
	set(connection, "name", "Johan")
	set(connection, "age", 42)
	set(connection, "vaccinated", true)
	fmt.Println()

	fmt.Println("Checking existence of values...")
	exists(connection, "name")
	exists(connection, "age")
	exists(connection, "vaccinated")
	exists(connection, "nonexistent")
	fmt.Println()

	fmt.Println("Getting values...")
	getString(connection, "name")
	getInt(connection, "age")
	getBool(connection, "vaccinated")
	fmt.Println()

	fmt.Println("Incrementing values...")
	// Can only increment int values
	//incr(connection, "name")
	incr(connection, "age")
	getInt(connection, "age")
	incrby(connection, "age", 5)
	getInt(connection, "age")
	fmt.Println()

	fmt.Println("Decrementing values...")
	// Can only decrement int values
	//decr(connection, "name")
	decr(connection, "age")
	getInt(connection, "age")
	decrby(connection, "age", 7)
	getInt(connection, "age")
	fmt.Println()

	fmt.Println("Expiring and persisting...")
	ttl(connection, "age")
	expire(connection, "age", 2)
	ttl(connection, "age")
	pause(1)
	//pause(3)
	persist(connection, "age")
	ttl(connection, "age")
	exists(connection, "age")
	fmt.Println()

	fmt.Println("Deleting value...")
	delete(connection, "vaccinated")
	exists(connection, "vaccinated")
	fmt.Println()

	fmt.Println("Lists...")
	rpush(connection, "names", "Bob")
	rpush(connection, "names", "Carol")
	lpush(connection, "names", "Alice")
	llen(connection, "names")
	lrange(connection, "names", 0, -1)
	lrange(connection, "names", 1, 2)
	rpop(connection, "names")
	lrange(connection, "names", 0, -1)
	lpop(connection, "names")
	lrange(connection, "names", 0, -1)
	fmt.Println()

	fmt.Println("Sets...")
	sadd(connection, "vegetarians", "Alice")
	sadd(connection, "vegetarians", "Carol")
	// Can add the same value multiple times - it's a set, so it's unique
	sadd(connection, "vegetarians", "Alice")
	sadd(connection, "clowns", "Bob")
	sadd(connection, "clowns", "Carol")
	sismember(connection, "vegetarians", "Alice")
	sismember(connection, "vegetarians", "Bob")
	sismember(connection, "vegetarians", "Carol")
	sismember(connection, "vegetarians", "Dave")
	sismember(connection, "clowns", "Alice")
	sismember(connection, "clowns", "Bob")
	sismember(connection, "clowns", "Carol")
	sismember(connection, "clowns", "Dave")
	smembers(connection, "vegetarians")
	smembers(connection, "clowns")
	smembers(connection, "nonexistent")
	srem(connection, "clowns", "Carol")
	smembers(connection, "clowns")
	// Popping from a set removes a random member...
	spop(connection, "vegetarians")
	smembers(connection, "vegetarians")
	// TODO: How can we do SUNION with redigo?
	fmt.Println()

	fmt.Println("Sorted sets...")
	zadd(connection, "ages", 42, "Alice")
	zadd(connection, "ages", 65, "Bob")
	zadd(connection, "ages", 21, "Carol")
	zrange(connection, "ages", 0, -1)
	zrange(connection, "ages", 1, 2)
	fmt.Println()

	fmt.Println("Hash maps...")
	hset(connection, "person:alice", "name", "Alice")
	hset(connection, "person:alice", "age", 42)
	hset(connection, "person:alice", "vegetarian", true)
	hmset(connection, "person:bob", map[string]interface{}{
		"name":       "Bob",
		"age":        65,
		"vegetarian": false})
	hget(connection, "person:alice", "name")
	hget(connection, "person:alice", "age")
	hget(connection, "person:alice", "vegetarian")
	hgetall(connection, "person:alice")
	hgetall(connection, "person:bob")
	hincrby(connection, "person:alice", "age", 5)
	hgetall(connection, "person:alice")
	hdel(connection, "person:alice", "vegetarian")
	hgetall(connection, "person:alice")
	fmt.Println()

	fmt.Println("Pipelining...")
	pipeline(connection)
	fmt.Println()

	// TODO: Pipelining using Send, Flush and Receive...
	// NOTE: The Do method combines the functionality of Send, Flush and Receive

	fmt.Println("Done.")
}

func cleanup(connection redigo.Conn) {
	keys, err := redis.Strings(connection.Do("KEYS", "*"))
	if err != nil {
		handleError("Error getting keys", err)
	}
	for _, key := range keys {
		fmt.Printf("Deleting key '%s'...\n", key)
		delete(connection, key)
	}
}

func set(connection redigo.Conn, key string, value interface{}) {
	fmt.Printf("Setting key '%s' to value: %v\n", key, value)
	_, err := connection.Do("SET", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error setting '%s'", key), err)
	}
}

func exists(connection redigo.Conn, key string) {
	exists, err := redis.Bool(connection.Do("EXISTS", key))
	if err != nil {
		handleError(fmt.Sprintf("Error checking if '%s' exists", key), err)
	}
	if exists {
		fmt.Printf("Key '%s' exists\n", key)
	} else {
		fmt.Printf("Key '%s' does not exist\n", key)
	}
}

func get(connection redigo.Conn, key string) (interface{}, error) {
	//fmt.Printf("Getting value for key '%s'...\n", key)
	value, err := connection.Do("GET", key)
	if err != nil {
		handleError(fmt.Sprintf("Error getting value for key '%s'", key), err)
	}
	//fmt.Printf("Returning value: %v\n", value)
	return value, err
}

func getString(connection redigo.Conn, key string) {
	value, err := redigo.String(get(connection, key))
	if err != nil {
		handleError("Error converting value to string", err)
	}
	fmt.Printf("Key '%s' has string value: %s\n", key, value)
}

func getInt(connection redigo.Conn, key string) {
	value, err := redigo.Int(get(connection, key))
	if err != nil {
		handleError("Error converting value to int", err)
	}
	fmt.Printf("Key '%s' has int value: %d\n", key, value)
}

func getBool(connection redigo.Conn, key string) {
	value, err := redigo.Bool(get(connection, key))
	if err != nil {
		handleError("Error converting value to bool", err)
	}
	fmt.Printf("Key '%s' has bool value: %t\n", key, value)
}

func delete(connection redigo.Conn, key string) {
	fmt.Printf("Deleting key '%s'...\n", key)
	_, err := connection.Do("DEL", key)
	if err != nil {
		handleError(fmt.Sprintf("Error deleting key '%s'", key), err)
	}
}

func incr(connection redigo.Conn, key string) {
	fmt.Printf("Incrementing key '%s'...\n", key)
	_, err := connection.Do("INCR", key)
	if err != nil {
		handleError(fmt.Sprintf("Error incrementing key '%s'", key), err)
	}
}

func incrby(connection redigo.Conn, key string, value int) {
	fmt.Printf("Incrementing key '%s' by %d...\n", key, value)
	_, err := connection.Do("INCRBY", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error incrementing key '%s' by %d", key, value), err)
	}
}

func decr(connection redigo.Conn, key string) {
	fmt.Printf("Decrementing key '%s'...\n", key)
	_, err := connection.Do("DECR", key)
	if err != nil {
		handleError(fmt.Sprintf("Error decrementing key '%s'", key), err)
	}
}

func decrby(connection redigo.Conn, key string, value int) {
	fmt.Printf("Decrementing key '%s' by %d...\n", key, value)
	_, err := connection.Do("DECRBY", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error decrementing key '%s' by %d", key, value), err)
	}
}

func expire(connection redigo.Conn, key string, seconds int) {
	fmt.Printf("Setting key '%s' to expire in %d seconds...\n", key, seconds)
	_, err := connection.Do("EXPIRE", key, seconds)
	if err != nil {
		handleError(fmt.Sprintf("Error setting key '%s' to expire in %d seconds", key, seconds), err)
	}
}

func ttl(connection redigo.Conn, key string) {
	ttl, err := redigo.Int(connection.Do("TTL", key))
	if err != nil {
		handleError(fmt.Sprintf("Error getting TTL for key '%s'", key), err)
	}
	fmt.Printf("TTL for key '%s' is %d\n", key, ttl)
}

func persist(connection redigo.Conn, key string) {
	fmt.Printf("Removing expiration for key '%s'...\n", key)
	_, err := connection.Do("PERSIST", key)
	if err != nil {
		handleError(fmt.Sprintf("Error removing expiration for key '%s'", key), err)
	}
}

func rpush(connection redigo.Conn, key string, value interface{}) {
	fmt.Printf("Appending value '%v' to list '%s'...\n", value, key)
	_, err := connection.Do("RPUSH", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error pushing value '%v' to list '%s'", value, key), err)
	}
}

func lpush(connection redigo.Conn, key string, value interface{}) {
	fmt.Printf("Prepending value '%v' to list '%s'...\n", value, key)
	_, err := connection.Do("LPUSH", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error pushing value '%v' to list '%s'", value, key), err)
	}
}

func llen(connection redigo.Conn, key string) {
	length, err := redigo.Int(connection.Do("LLEN", key))
	if err != nil {
		handleError(fmt.Sprintf("Error getting length of list '%s'", key), err)
	}
	fmt.Printf("Length of list '%s': %d\n", key, length)
}

func lrange(connection redigo.Conn, key string, start, stop int) {
	values, err := redigo.Strings(connection.Do("LRANGE", key, start, stop))
	if err != nil {
		handleError(fmt.Sprintf("Error getting values from list '%s'", key), err)
	}
	if start == 0 && stop == -1 {
		fmt.Printf("Values in list '%s': %v\n", key, values)
	} else {
		fmt.Printf("Values in list '%s[%d:%d]': %v\n", key, start, stop, values)
	}
}

func rpop(connection redigo.Conn, key string) {
	value, err := redigo.String(connection.Do("RPOP", key))
	if err != nil {
		handleError(fmt.Sprintf("Error popping value from list '%s'", key), err)
	}
	fmt.Printf("Rpopped value: %s\n", value)
}

func lpop(connection redigo.Conn, key string) {
	value, err := redigo.String(connection.Do("LPOP", key))
	if err != nil {
		handleError(fmt.Sprintf("Error popping value from list '%s'", key), err)
	}
	fmt.Printf("Lpopped value: %s\n", value)
}

func sadd(connection redigo.Conn, key string, value interface{}) {
	fmt.Printf("Adding value '%v' to set '%s'...\n", value, key)
	_, err := connection.Do("SADD", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error adding value '%v' to set '%s'", value, key), err)
	}
}

func sismember(connection redigo.Conn, key string, value interface{}) {
	exists, err := redigo.Bool(connection.Do("SISMEMBER", key, value))
	if err != nil {
		handleError(fmt.Sprintf("Error checking if value '%v' is in set '%s'", value, key), err)
	}
	if exists {
		fmt.Printf("'%v' is a member of set '%s'\n", value, key)
	} else {
		fmt.Printf("'%v' is NOT a member of set '%s'\n", value, key)
	}
}

func smembers(connection redigo.Conn, key string) {
	values, err := redigo.Strings(connection.Do("SMEMBERS", key))
	if err != nil {
		handleError(fmt.Sprintf("Error getting values from set '%s'", key), err)
	}
	fmt.Printf("Members of set '%s': %v\n", key, values)
}

func srem(connection redigo.Conn, key string, value interface{}) {
	fmt.Printf("Removing member '%v' from set '%s'...\n", value, key)
	_, err := connection.Do("SREM", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error removing value '%v' from set '%s'", value, key), err)
	}
}

func spop(connection redigo.Conn, key string) {
	value, err := redigo.String(connection.Do("SPOP", key))
	if err != nil {
		handleError(fmt.Sprintf("Error popping value from set '%s'", key), err)
	}
	fmt.Printf("Spopped value: %s\n", value)
}

func zadd(connection redigo.Conn, key string, score int, member interface{}) {
	fmt.Printf("Adding member '%v' with score %d to sorted set '%s'...\n", member, score, key)
	_, err := connection.Do("ZADD", key, score, member)
	if err != nil {
		handleError(fmt.Sprintf("Error adding member '%v' with score %d to sorted set '%s'", member, score, key), err)
	}
}

func zrange(connection redigo.Conn, key string, start, stop int) {
	values, err := redigo.Strings(connection.Do("ZRANGE", key, start, stop))
	if err != nil {
		handleError(fmt.Sprintf("Error getting values from sorted set '%s'", key), err)
	}
	if start == 0 && stop == -1 {
		fmt.Printf("Members of sorted set '%s': %v\n", key, values)
	} else {
		fmt.Printf("Members of sorted set '%s[%d:%d]': %v\n", key, start, stop, values)
	}
}

func hset(connection redigo.Conn, key string, field string, value interface{}) {
	fmt.Printf("Setting field '%s' to value '%v' in hash map '%s'...\n", field, value, key)
	_, err := connection.Do("HSET", key, field, value)
	if err != nil {
		handleError(fmt.Sprintf("Error setting field '%s' to value '%v' in hash map '%s'", field, value, key), err)
	}
}

func hmset(connection redigo.Conn, key string, values map[string]interface{}) {
	fmt.Printf("Setting multiple fields in hash map '%s'...\n", key)
	args := []interface{}{key}
	for field, value := range values {
		args = append(args, field, value)
	}
	_, err := connection.Do("HMSET", args...)
	if err != nil {
		handleError(fmt.Sprintf("Error setting multiple fields in hash map '%s'", key), err)
	}
}

func hget(connection redigo.Conn, key string, field string) {
	value, err := redigo.String(connection.Do("HGET", key, field))
	if err != nil {
		handleError(fmt.Sprintf("Error getting field '%s' from hash map '%s'", field, key), err)
	}
	fmt.Printf("Value of field '%s' in hash map '%s': %v\n", field, key, value)
}

func hgetall(connection redigo.Conn, key string) {
	values, err := redigo.Strings(connection.Do("HGETALL", key))
	if err != nil {
		handleError(fmt.Sprintf("Error getting values from hash map '%s'", key), err)
	}
	fmt.Printf("Values in hash map '%s': %v\n", key, values)
}

func hincrby(connection redigo.Conn, key string, field string, value int) {
	fmt.Printf("Incrementing field '%s' in hash map '%s' by %d...\n", field, key, value)
	_, err := connection.Do("HINCRBY", key, field, value)
	if err != nil {
		handleError(fmt.Sprintf("Error incrementing field '%s' in hash map '%s' by %d", field, key, value), err)
	}
}

func hdel(connection redigo.Conn, key string, field string) {
	fmt.Printf("Deleting field '%s' from hash map '%s'...\n", field, key)
	_, err := connection.Do("HDEL", key, field)
	if err != nil {
		handleError(fmt.Sprintf("Error deleting field '%s' from hash map '%s'", field, key), err)
	}
}

func pipeline(connection redigo.Conn) {
	// Send 3 SET commands...
	if err := connection.Send("SET", "firstname", "Johan"); err != nil {
		handleError("Error sending command", err)
	}
	if err := connection.Send("SET", "middlename", "Frederik"); err != nil {
		handleError("Error sending command", err)
	}
	if err := connection.Send("SET", "surname", "Steyn"); err != nil {
		handleError("Error sending command", err)
	}
	// Send 3 GET commands...
	if err := connection.Send("GET", "firstname"); err != nil {
		handleError("Error sending command", err)
	}
	if err := connection.Send("GET", "middlename"); err != nil {
		handleError("Error sending command", err)
	}
	if err := connection.Send("GET", "surname"); err != nil {
		handleError("Error sending command", err)
	}
	// Flush the commands...
	if err := connection.Flush(); err != nil {
		handleError("Error flushing", err)
	}
	// Receive the SET responses...
	if _, err := connection.Receive(); err != nil {
		handleError("Error receiving", err)
	}
	if _, err := connection.Receive(); err != nil {
		handleError("Error receiving", err)
	}
	if _, err := connection.Receive(); err != nil {
		handleError("Error receiving", err)
	}
	// Receive the GET responses...
	var err error
	var firstname string
	var middlename string
	var surname string
	//if firstname, err = connection.Receive(); err != nil {
	if firstname, err = redigo.String(connection.Receive()); err != nil {
		handleError("Error receiving", err)
	}
	fmt.Printf("Firstname: %s\n", firstname)
	if middlename, err = redigo.String(connection.Receive()); err != nil {
		handleError("Error receiving", err)
	}
	fmt.Printf("Middlename: %s\n", middlename)
	if surname, err = redigo.String(connection.Receive()); err != nil {
		handleError("Error receiving", err)
	}
	fmt.Printf("Surname: %s\n", surname)

}

func pause(seconds int) {
	fmt.Printf("Sleeping for %d seconds...\n", seconds)
	time.Sleep(time.Second * time.Duration(seconds))
}

func handleError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
	os.Exit(1)
}
