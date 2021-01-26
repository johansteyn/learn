#!/bin/bash

command=`basename $0`

function usage {
  echo ""
  echo "USAGE: $command <yyyy-MM-dd> <password> [<limit>]"  
  echo ""
}

if [ $# -ne 2 ] && [ $# -ne 3 ]
then
  usage
  exit
fi

dt=$1
password=$2
if [ $# -eq 3 ]
then
  limit=$3
fi

echo "Running Scala ETL application..."
scala -Dcom.datastax.driver.NATIVE_TRANSPORT_MAX_FRAME_SIZE_IN_MB=512 target/scala-2.12/etl-scala-assembly-1.0.jar $dt $password $limit

# NOTES
# The NATIVE_TRANSPORT_MAX_FRAME_SIZE_IN_MB definition is needed to get passed this exception that occurs after reading 50K rows:
#   Exception: com.datastax.driver.core.exceptions.FrameTooLongException: Response frame exceeded maximum allowed length

