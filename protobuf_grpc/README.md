# Protobuff GRPC - talk

## Installation

```bash
docker run --rm -it golang
```

``` bash
apt update; apt install vim unzip -y
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.17.0/protoc-3.17.0-linux-x86_64.zip
unzip protoc-3.17.0-linux-x86_64.zip
mv ./bin/protoc /usr/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
mkdir ~/demo ; cd ~/demo
```

### demo proto

```bash
vim ./demo.proto
```

```proto3
syntax = "proto3";

package demo;
option go_package="github.com/wr4thon/demo";

message Target {
	string Name = 1;
}

message Void {}

service Greeter {
	rpc Greet(Target) returns (Void);
}
```
### generation
```bash
protoc --go_out=. --go-grpc_out=. ./demo.proto
```

## References

Resources I used:

- [justforfunc #30: The Basics of Protocol Buffers](https://www.youtube.com/watch?v=_jQ3i_fyqGA)
- [justforfunc #31: gRPC Basics](https://www.youtube.com/watch?v=uolTUtioIrc)
- [google developers - protocol buffers](https://developers.google.com/protocol-buffers)
- [Wikipedia - remote procedure call](https://en.wikipedia.org/wiki/Remote_procedure_call)
- [grpc.io](https://grpc.io/)
- [Protocol Buffers Crash Course](https://www.youtube.com/watch?v=46O73On0gyI)
- [protobuf github](https://github.com/protocolbuffers/protobuf)

Links I found helpful during the preparation:

- [proto3 documentation](https://developers.google.com/protocol-buffers/docs/proto3)
- [mustEmbedUnimplemented*** method appear in grpc-server](https://github.com/grpc/grpc-go/issues/3794)
