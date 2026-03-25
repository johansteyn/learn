#!/bin/bash

for i in stdout-*; do echo "------------------------"; echo $i; tail $i; done


