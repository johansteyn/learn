package main

import (
	"fmt"
	"os"

	redigo "github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("Go 3rd-party Library: redigo")
	fmt.Println()

	address := "redis-14240.c328.europe-west3-1.gce.cloud.redislabs.com:14240"
	password := "isYucMJRfpVKkNKrONeiYV0qDd42JDk5"
	connection, err := redigo.Dial("tcp", address, redigo.DialPassword(password))
	if err != nil {
		handleError("Error connecting to Redis", err)
	}
	defer connection.Close()
	fmt.Printf("Connected to redis: %+v\n", connection)
	fmt.Println()

	fmt.Println("Setting values...")
	set(connection, "name", "Johan")
	set(connection, "age", 42)
	fmt.Println()

	fmt.Println("Getting values...")
	key := "name"
	strValue, err := redigo.String(get(connection, key))
	if err != nil {
		handleError("Error converting value to string", err)
	}
	fmt.Printf("Key '%s' has value: %s\n", key, strValue)
	key = "age"
	intValue, err := redigo.Int(get(connection, key))
	if err != nil {
		handleError("Error converting value to int", err)
	}
	fmt.Printf("Key '%s' has value: %d\n", key, intValue)
	fmt.Println()

	fmt.Println("Done.")
}

func set(connection redigo.Conn, key string, value interface{}) {
	fmt.Printf("Setting key '%s' to value: %v\n", key, value)
	_, err := connection.Do("SET", key, value)
	if err != nil {
		handleError(fmt.Sprintf("Error setting '%s'", key), err)
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

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
