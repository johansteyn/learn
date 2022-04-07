#!/bin/bash

# Example usages:
#   test.sh test-cli
#   test.sh test-udp-client

start=`date +%s`
for i in `seq 10`; do echo $i; $*; done
end=`date +%s`
duration=$((end-start))
echo "Time taken: $duration seconds"

