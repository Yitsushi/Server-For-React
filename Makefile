all: clear mac windows linux

mac:
	go build -o bin/mac-standaloneStaticServer

windows:
	GOOS=windows GOARCH=386 go build -o bin/windows-standaloneStaticServer-i386.bin
	GOOS=windows GOARCH=amd64 go build -o bin/windows-standaloneStaticServer-amd64.bin

linux:
	GOOS=linux GOARCH=386 go build -o bin/linux-standaloneStaticServer-i386.bin
	GOOS=linux GOARCH=amd64 go build -o bin/linux-standaloneStaticServer-amd64.bin

clear:
	mkdir -p bin
	rm -f bin/*
