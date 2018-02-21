#!/bin/bash

# MacOS needs gdate for %N to work...
# $ brew install coreutils
alias date='gdate'
STARTTIME=`gdate +%s.%N`
java -cp target/wordcounter-1.0.jar SimpleWordCounter
ENDTIME=`gdate +%s.%N`
TIMEDIFF=`echo "$ENDTIME - $STARTTIME" | bc | awk -F"." '{print $1"."substr($2,1,3)}'`
echo "$TIMEDIFF"

