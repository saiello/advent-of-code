#!/bin/bash

set -xeu

source .env
day=$(date '+%d')
mkdir -p 2023/day${day}
curl -sq -H "cookie: ${aoc_cookie}"  https://adventofcode.com/2023/day/${day#0}/input -o 2023/day${day}/input.txt
cat 2023/day${day}/input.txt