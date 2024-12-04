#!/bin/bash
DAY=$1
STEP=$2

go run "days/${DAY}/${STEP}/main.go" < "days/${DAY}/${STEP}/input.txt"
