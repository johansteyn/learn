#!/bin/bash

STARTTIME=`date +%s.%N`
java -jar target/wordcounter-1.0.jar
ENDTIME=`date +%s.%N`
TIMEDIFF=`echo "$ENDTIME - $STARTTIME" | bc | awk -F"." '{print $1"."substr($2,1,3)}'`
echo "$TIMEDIFF"

