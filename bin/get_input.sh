#!/bin/bash

set -xeu

source .env
today=$(date '+%d')
day=${1:-${today}}
year=$(date '+%Y')
mkdir -p ${year}/day${day}
curl -sq --cookie "session=${aoc_cookie}"  https://adventofcode.com/${year}/day/${day#0}/input -o ${year}/day${day}/input.txt
cat ${year}/day${day}/input.txt