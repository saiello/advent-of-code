My solutions of the advent of code using Golang, to play with it. 


Some alias to run solution of the current day against the input and the example. ( since 2024 )

```
alias aoi='cat `date +"%Y/day%d/input.txt"` | go run `date +"%Y/day%d/solve.go"`'                                
alias aoe='cat `date +"%Y/day%d/example1.txt"` | go run `date +"%Y/day%d/solve.go"`' 
```

A script to prepare the daily folder and collect the input and example file.
```
./bin/get_input.sh
```

> [!Warning] Check the example
> The fetched example might be not exact.

