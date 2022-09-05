# Install tools/packages

윈도우 WSL2 환경에서 실습을 하였으며, 각 장에서 필요한 툴과 패키지들의 설치를 정리해 두었다. 
쉘 명령어들은 바로 복사해서 쓸 수 있도록 프롬프트를 생략해두었다. 

## 1장. 프로젝트 시작

Go 설치: https://go.dev/dl/

## 2장. 프로토콜 버퍼와 구조체

### protoc

참고: https://grpc.io/docs/protoc-installation/

```zsh
# 설치 후 권한 이슈가 있어 홈 디렉토리에 설치하도록 함
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

## 3장. 로그 패키지 작성

```
go get -u github.com/stretchr/testify
go get -u github.com/tysonmote/gommap

```

## 4장. gRPC 요청 처리 

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

```

## 5장. 서비스 보안

```
go install github.com/cloudflare/cfssl/cmd/cfssl@v1.6.1
go install github.com/cloudflare/cfssl/cmd/cfssljson@v1.6.1

go get github.com/casbin/casbin@v1.9.1

go get github.com/grpc-ecosystem/go-grpc-middleware
go get github.com/grpc-ecosystem/go-grpc-middleware/auth
```

## 6장. 시스템 관측

```
go get go.uber.org/zap@v1.21.0
go get go.opencensus.io@v0.23.0
```

## 7장. 서버 간 서비스 디스커버리

```
go get github.com/hashicorp/serf@v0.9.7
go get github.com/travisjeffery/go-dynaport@v1.0.0
```

## 8장. 합의를 통한 서비스 간 조율

```
go get github.com/hashicorp/raft@v1.3.6
go mod edit -replace github.com/hashicorp/raft-boltdb=github.com/travisjeffery/raft-boltdb@v1.0.0
go get -u github.com/soheilhy/cmux
```

## 9장. 클라이언트 측 서버 디스커버리와 로드밸런스

설치 없음

## 10장. 로컬에서 쿠버네티스로 배포

### kubectl

kubectl 설치 및 설정. 2장에서와 마찬가지로 `~/.local`에 설치한다.

```
curl -LO https://dl.k8s.io/release/v1.21.13/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p ~/.local/bin
mv ./kubectl ~/.local/bin/kubectl
echo 'export PATH="$PATH:$HOME/.local/bin"' >> ~/.zshenv
source ~/.zshenv
```

### Go 패키지와 도구 
```
go install sigs.k8s.io/kind@v0.12.0
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

### Helm
```
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

```

## 11장. 클라우드에서 쿠버네티스로 배포

### gcloud SDK 설치
```
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-383.0.1-linux-x86_64.tar.gz 
tar -zxvf google-cloud-cli-383.0.1-linux-x86_64.tar.gz  -C ~/
~/google-cloud-sdk/install.sh
source ~/.zshrc
gcloud --version
rm -rf google-cloud-cli-383.0.1-linux-x86_64.tar.gz 
```
