GOROOT=/Users/svdu/go/go1.13.1
GOPATH=/Users/svdu/go
VERSION=0.01.5


GOOS=linux
GOARCH=arm
GOARM=5
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/jan_modbus_tcp/bin/$VERSION/arm/jan_modbus_tcp /Users/svdu/GolandProjects/jan_modbus_tcp/main.go
GOOS=windows
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/jan_modbus_tcp/bin/$VERSION/windows/jan_modbus_tcp /Users/svdu/GolandProjects/jan_modbus_tcp/main.go
GOOS=linux
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/jan_modbus_tcp/bin/$VERSION/linux/jan_modbus_tcp /Users/svdu/GolandProjects/jan_modbus_tcp/main.go
GOOS=darwin
GOARCH=amd64
$GOROOT/bin/go build -o /Users/svdu/GolandProjects/jan_modbus_tcp/bin/$VERSION/macos/jan_modbus_tcp /Users/svdu/GolandProjects/jan_modbus_tcp/main.go