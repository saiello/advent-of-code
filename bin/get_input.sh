#!/bin/bash

set -xeu

source .env
today=$(date '+%d')
day=${1:-${today}}
year=$(date '+%Y')
mkdir -p ${year}/day${day}
curl -sq --cookie "session=${aoc_cookie}"  https://adventofcode.com/${year}/day/${day#0}/input -o ${year}/day${day}/input.txt

echo "Input"
head ${year}/day${day}/input.txt

curl -s https://adventofcode.com/${year}/day/${day#0} | xmllint --nowarning --html --xpath '/html/body/main/article/pre/code/text()' - > ${year}/day${day}/example1.txt

echo "Example"
cat ${year}/day${day}/example1.txt