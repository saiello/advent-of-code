My solutions of the advent of code 


Set the aoc cookie in a .env file

```
echo "export aoc_cookie=<session_aoc_cookie>" > .env
```

```
source .env
day=$(date '+%d')
curl -H 'cookie: session=${aoc_session};'  https://adventofcode.com/2023/day/${day#0}/input -o 2023/day${day}/input.txt
```