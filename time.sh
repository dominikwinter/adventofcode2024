#!/bin/bash
DAY=$1
STEP=$2
EX=$3

go build -ldflags="-s -w" -trimpath -o aoc2024

# warmup
./aoc2024 "${DAY}${STEP}" "input/${DAY}${EX}.txt" > /dev/null

time ./aoc2024 "${DAY}${STEP}" "input/${DAY}${EX}.txt" > /dev/null

rm ./aoc2024
