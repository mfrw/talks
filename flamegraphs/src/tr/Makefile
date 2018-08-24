all:
	go build simple0.go
	go build simple1.go
	go build simple2.go
.PHONY: clean test
test:
	@echo "Running simple0"
	@./simple0 -trace 0.out
	@echo "Running simple1"
	@./simple1 -trace 1.out
	@echo "Running simple2"
	@./simple2 -trace 2.out
clean:
	rm -f *.out simple0 simple1 simple2
	
