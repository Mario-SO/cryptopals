.PHONY: test set-1 set-2 set-3 set-4 set-5 set-6 set-7 set-8

# Run all tests
test:
	go test ./... -v

# Watch specific sets
set-1:
	air -- "set_1"

set-2:
	air -- "set_2"

set-3:
	air -- "set_3"

set-4:
	air -- "set_4"

set-5:
	air -- "set_5"

set-6:
	air -- "set_6"

set-7:
	air -- "set_7"

set-8:
	air -- "set_8"
