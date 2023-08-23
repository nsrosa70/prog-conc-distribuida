@ECHO OFF
set GO111MODULE=on
set GOPATH=C:\Users\user\go
set GOROOT=c:\Program Files\Go
set PATH=%PATH%;c:\Program Files\Go\bin;C:\Program Files\Git\bin;C:\Users\user\go\bin;C:\Program Files\protoc;C:\Users\user\go\pkg\mod\github.com\golang\protobuf@v1.5.2\protoc-gen-go
set PATH=%PATH%;C:\Users\user\go\pkg\mod\google.golang.org\grpc\cmd\protoc-gen-go-grpc@v1.2.0
set PATH=%PATH%;protoc-gen-go-grpc@v1.2
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative proto/calculadora.proto

cd C:\Users\user\go\prog-conc-distribuida\distribuida\fibonacci
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative proto/fibonacci.proto

cd C:\Users\user\go\prog-conc-distribuida\distribuida\calculadora\grpc
