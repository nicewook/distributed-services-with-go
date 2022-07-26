# Install tools

Install tools on WSL2

## Chapter 1

Install Go: 

## Chapter 2

### protoc

- reference: https://grpc.io/docs/protoc-installation/

```zsh
# root issue - so change it to ~
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip
mkdir -p $HOME/.local/protobuf
unzip protoc-3.19.4-linux-x86_64.zip -d $HOME/.local/protobuf
rm ./protoc-3.19.4-linux-x86_64.zip
chmod +x $HOME/.local/protobuf/bin/protoc
echo 'export PATH="$PATH:$HOME/.local/protobuf/bin"' >> ~/.zshenv
protoc --version
```

### protobuf runtime

Go로 컴파일하기 위한 런타임 설치 

```
go install google.golang.org/protobuf/...@v1.27.1
```

## Chapter 3

```
go get -u github.com/stretchr/testify
go get -u github.com/tysonmote/gommap

```

## Chapter 4

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

```

## Chapter 5

```
go install github.com/cloudflare/cfssl/cmd/cfssl@v1.6.1
go install github.com/cloudflare/cfssl/cmd/cfssljson@v1.6.1

go get github.com/casbin/casbin@v1.9.1

go get github.com/grpc-ecosystem/go-grpc-middleware
go get github.com/grpc-ecosystem/go-grpc-middleware/auth
```

## Chapter 6

```
go get go.uber.org/zap@v1.21.0
go get go.opencensus.io@v0.23.0
```

## Chapter 7

```
go get github.com/hashicorp/serf@v0.9.7
go get github.com/travisjeffery/go-dynaport@v1.0.0
```

## Chapter 8

```
go get github.com/hashicorp/raft@v1.3.6
go mod edit -replace github.com/hashicorp/raft-boltdb=github.com/travisjeffery/raft-boltdb@v1.0.0
```