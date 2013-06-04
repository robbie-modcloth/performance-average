#!/bin/bash

DATA=/Users/r.mckinstry/workspace/Data

echo "Formatting the code"
go fmt
echo "The code has been formatted"

echo "Compiling the code"
go build
echo "The code has been compiled"

echo "Executing the code"
./performance-average $DATA
echo "The code has been executed"

rm performance-average
