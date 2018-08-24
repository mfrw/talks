all:
	go build simple0.go
	go build simple1.go
	go build simple2.go
	go build simple3.go
.PHONY: clean test
test:
	@echo "Running simple0"
	@./simple0 -profile 0.out -out 0.txt
	@echo "Running simple1"
	@./simple1 -profile 1.out -out 1.txt
	@echo "Running simple2"
	@./simple2 -profile 2.out -out 2.txt
	@echo "Running simple3"
	@./simple2 -profile 3.out -out 3.txt
clean:
	rm -f *.out simple0 simple1 simple2 simple3 0.txt 1.txt 2.txt 3.txt
	
