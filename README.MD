
# chat go with  grpc

before install depend
>set path  
"export PATH=$PATH:$HOME/go/bin"

instal latest protoc for golang
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

instal protoc
(for linux)
> sudo apt-get install protoc 

## for generate protobuf files do
###  protoc  --go_out=chat_pb chat_pb/chat.proto



