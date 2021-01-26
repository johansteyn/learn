#!/bin/bash

command=`basename $0`

function usage {
  echo ""
  echo "USAGE: $command <password> [<limit>]"
  echo ""
}

if [ $# -ne 1 ] && [ $# -ne 2 ]
then
  usage
  exit
fi

password=$1
if [ $# -eq 2 ]
then
  limit=$2
fi

echo "Running Scala Cassandra application..."
scala -Dcom.datastax.driver.NATIVE_TRANSPORT_MAX_FRAME_SIZE_IN_MB=512 target/scala-2.12/cassandra-scala-assembly-1.0.jar $password $limit

# NOTES
# The NATIVE_TRANSPORT_MAX_FRAME_SIZE_IN_MB definition is needed to get passed this exception that occurs after reading 50K rows:
#   Exception: com.datastax.driver.core.exceptions.FrameTooLongException: Response frame exceeded maximum allowed length

