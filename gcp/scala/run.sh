#!/bin/bash

command=`basename $0`

function usage {
  echo ""
  echo "USAGE: $command <bucket>"
  echo ""
}

if [ $# -ne 1 ]
then
  usage
  exit
fi

bucket=$1
export GOOGLE_APPLICATION_CREDENTIALS=wmt-1f780b38bd7b0384e53292de20.json
echo "Running Scala GoogleStorage application for '$bucket' on GCP..."
scala target/scala-2.12/gcp-scala-assembly-1.0.jar $bucket

