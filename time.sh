#!/bin/bash

set -e

YEAR=$1
DAY=$2
STEP=$3

BIN="${YEAR}_${DAY}_${STEP}"

go build -ldflags="-s -w" -trimpath -o "${BIN}" "${YEAR}/${DAY}/${STEP}/main.go"

# warm up
"./${BIN}" < "${YEAR}/${DAY}/${STEP}/input.txt" > /dev/null

time "./${BIN}" < "${YEAR}/${DAY}/${STEP}/input.txt"

rm -f "${BIN}"
