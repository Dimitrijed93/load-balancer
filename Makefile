.PHONY : clean build run start execute debug dlv

clean:
	rm -r cmd/main

build:
	go build -o cmd/main

execute:
	./cmd/main

dlv:
	dlv debug ./main.go

run: clean build execute

debug: clean build dlv