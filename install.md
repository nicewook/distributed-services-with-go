# Install tools

Install tools on WSL2

## Chapter 1

Install Go: 

## Chapter 2

### protoc

```zsh
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip
sudo unzip protoc-3.19.4-linux-x86_64.zip -d /usr/local/protobuf
rm ./protoc-3.19.4-linux-x86_64.zip
sudo chmod +x /usr/local/protobuf/bin/protoc
echo 'export PATH="$PATH:/usr/local/protobuf/bin"' >> ~/.zshenv
protoc --version
```

### protobuf runtime

Go로 컴파일하기 위한 런타임 설치 

```
go install google.golang.org/protobuf/...@v1.27.1
```