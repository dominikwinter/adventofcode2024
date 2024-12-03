#!/bin/bash
DAY=$1
STEP=$2
EX=$3

go run main.go "${DAY}${STEP}" "input/${DAY}${EX}.txt"
