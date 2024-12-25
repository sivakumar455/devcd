.PHONY: clean build run

build:
	@echo "building devcd"
	go build -o devc ./devcd-ext/cmd

run: build 
	./devc -h

clean:
	rm ./devc
