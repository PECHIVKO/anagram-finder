# anagram-finder

_*How to run?*_

Clone repository and run command in terminal "go run main.go" OR "go build" "go run ./"

_*How to use?*_

1. Open new tab in terminal
2. Use command: " curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]' " to load dictionary
3. Use command " curl 'localhost:8080/get?word=foobar' " to get anagrams for words
