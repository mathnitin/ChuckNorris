#!/bin/bash

######
# Continuous 100 requests, back to back. 
#####

i="0"

while [ $i -lt 100 ]
do
    echo 'Calling curl:'
    curl http://localhost:5000
    echo 
done
