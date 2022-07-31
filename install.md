# Install tools/packages

윈도우 WSL2 환경에서 실습을 하였으며

각 장에서 필요한 툴과 패키지들의 설치를 정리해 두었다. 

## Chapter 1

Install Go: 

## Chapter 2

### protoc

- reference: https://grpc.io/docs/protoc-installation/

```zsh
# root issue - so change it to ~
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip
unzip protoc-3.19.4-linux-x86_64.zip -d ~/.local/protobuf
rm ./protoc-3.19.4-linux-x86_64.zip
echo 'export PATH="$PATH:$HOME/.local/protobuf/bin"' >> ~/.zshenv
source ~/.zshenv
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
go get -u github.com/soheilhy/cmux
```

## Chapter 10

### kubectl

kubectl 설치 및 실행할 수 있도록 설정

2장에서와 마찬가지로 ~/.local에 설치함

```
curl -LO https://dl.k8s.io/release/v1.21.13/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p ~/.local/bin
mv ./kubectl ~/.local/bin/kubectl
echo 'export PATH="$PATH:$HOME/.local/bin"' >> ~/.zshenv
source ~/.zshenv
```

```
go install sigs.k8s.io/kind@v0.12.0
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

```
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

```

### helm 설치 프로세스 정리(TODO)

## Chapter 11

### gcloud SDK 설치
```
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-383.0.1-linux-x86_64.tar.gz 
tar -zxvf google-cloud-cli-383.0.1-linux-x86_64.tar.gz  -C ~/
~/google-cloud-sdk/install.sh
source ~/.zshrc
gcloud --version
rm -rf google-cloud-cli-383.0.1-linux-x86_64.tar.gz 
```
### 기타 실습 정리하기(TODO)