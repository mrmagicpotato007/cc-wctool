#!/bin/bash
go build
echo "Running No of bytes command"
./wctool -c test.txt 
echo "Running No of words command"
./wctool -w test.txt
echo "Running No of chars command"
./wctool -m test.txt
echo "Running default command"
./wctool test.txt
