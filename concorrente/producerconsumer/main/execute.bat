@echo off
set GO111MODULE=off
set GOPATH=C:\Users\user\go\prog-conc-distribuida
set GOROOT=C:\Program Files\go

cd C:\Users\user\go\prog-conc-distribuida\concorrente\producerconsumer\main
go build -o server.exe main.go
rem go run C:\Users\user\go\prog-conc-distribuida\concorrente\producerconsumer\main\main.go
cd C:\Users\user\go\prog-conc-distribuida\concorrente\producerconsumer\main