#!/bin/bash
YEAR=$1
DAY=$2
STEP=$3

go run "${YEAR}/${DAY}/${STEP}/main.go" < "${YEAR}/${DAY}/${STEP}/input.txt"
