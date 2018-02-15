#!/bin/bash

run() {
	echo "======= $1 ======="
	java -cp target/fibonacci-1.0.jar Fibonacci $1
	java -cp target/fibonacci-1.0.jar Fibonacci $1
	java -cp target/fibonacci-1.0.jar Fibonacci $1
	echo "------------"
	java -cp target/fibonacci-1.0.jar MultiFibonacci $1 4
	java -cp target/fibonacci-1.0.jar MultiFibonacci $1 4
	java -cp target/fibonacci-1.0.jar MultiFibonacci $1 4
	echo "------------"
	java -jar target/fibonacci-1.0.jar $1
	java -jar target/fibonacci-1.0.jar $1
	java -jar target/fibonacci-1.0.jar $1
}

run 1000
run 2000
run 3000
run 4000
#run 5000
#run 6000
#run 7000
#run 8000
#run 9000
#run 10000

