#!/bin/bash

command=`basename $0`

function usage {
  echo ""
  echo "USAGE: $command <file>"
  echo ""
}

if [ $# -ne 1 ]
then
  usage
  exit
fi

file=$1
echo "Running Java application to parse '$file' using GSon..."
java -jar target/gson-java.jar $file

