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

echo "Running Scala application to parse '$file' using GSon..."
scala target/scala-2.12/gson-scala-assembly-1.0.jar $file

