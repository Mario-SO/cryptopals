.PHONY: test set-1 set-2 set-3 set-4 set-5 set-6 set-7 set-8

# Run all tests
test:
	go test ./... -v

# Watch specific sets
set-1:
	SET_NUMBER=set_1 air

set-2:
	SET_NUMBER=set_2 air

set-3:
	SET_NUMBER=set_3 air

set-4:
	SET_NUMBER=set_4 air

set-5:
	SET_NUMBER=set_5 air

set-6:
	SET_NUMBER=set_6 air

set-7:
	SET_NUMBER=set_7 air

set-8:
	SET_NUMBER=set_8 air
