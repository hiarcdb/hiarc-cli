build:
	go build -o hiarc main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=386 go build -o bin/darwin/hiarc main.go
	GOOS=freebsd GOARCH=386 go build -o bin/freebsd/hiarc main.go
	GOOS=linux GOARCH=386 go build -o bin/linux/hiarc main.go
	GOOS=windows GOARCH=386 go build -o bin/windows/hiarc main.go

run:
	go run main.go