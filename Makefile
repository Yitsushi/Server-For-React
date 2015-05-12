all: clear mac windows linux

mac:
	go build -o bin/mac-standaloneStaticServer standaloneStaticServer/main.go

windows:
	GOOS=windows GOARCH=386 go build -o bin/windows-standaloneStaticServer-i386.exe standaloneStaticServer/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows-standaloneStaticServer-amd64.exe standaloneStaticServer/main.go

linux:
	GOOS=linux GOARCH=386 go build -o bin/linux-standaloneStaticServer-i386.bin standaloneStaticServer/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux-standaloneStaticServer-amd64.bin standaloneStaticServer/main.go

clear:
	mkdir -p bin
	rm -f bin/*
