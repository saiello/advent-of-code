#!/bin/bash

set -xeu

source .env
today=$(date '+%d')
day=${1:-${today}}
year=$(date '+%Y')

daily_folder=${year}/day${day}


mkdir -p ${daily_folder}

# get input
curl -sq --cookie "session=${aoc_cookie}"  https://adventofcode.com/${year}/day/${day#0}/input -o ${daily_folder}/input.txt

echo "Input"
head ${daily_folder}/input.txt

# Get example
curl -s https://adventofcode.com/${year}/day/${day#0} | xmllint --nowarning --html --xpath '/html/body/main/article/pre/code/text()' - > ${daily_folder}/example1.txt

echo "Example"
cat ${daily_folder}/example1.txt

cp .template/solve.go ${daily_folder}/