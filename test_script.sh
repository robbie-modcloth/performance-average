#!/bin/bash

#personalize this script with the location of your own data
DATA=/Users/r.mckinstry/workspace/Data
TESTDATA=/Users/r.mckinstry/workspace/TEST

echo "Formatting the code"
go fmt
echo "The code has been formatted"

echo "Compiling the code"
go build
echo "The code has been compiled"

echo "Executing the code"
./performance-average $TESTDATA
./performance-average $DATA
echo "The code has been executed"

rm performance-average
