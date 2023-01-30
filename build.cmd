del /Q bin
set GOOS=linux
set GOARCH = amd64
go build -ldflags="-s -w" -o bin/welcome endpoints/aws/welcome.go 