# Chapter 2. Structure Data With Protobuf

## ch1 요약

- Record가 있고, 그걸 담는 Log가 있다. 
- Log는 Append, Read 메서드로 추가하고, 읽는다. 
- httpServer는 Log를 가지고, Log 메서드를 이용해 처리하는 handleProduce, handleConsume 핸들러가 있다.
- NewHTTPServer를 통해 진짜 서버를 돌리는데 mux에서는 httpServer의 핸들러 메서드로 Log를 추가하고 읽어서 회신한다. 

## ch2 개요

1. `protoc` 컴파일러, protobuf Go 런타임을 설치한 다음
2. `log.proto`를 만들고,
3. `Makefile`을 만들어 둔 다음
4. 컴파일해서 `log.pb.go` 파일을 만들다. 

## ch2 준비사항

프로토콜 버퍼 컴파일러 설치

참고로 나는 경로를 `~/zshrc`에 넣어주었다. (아래에는 ~/.zshenv)
```
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-osx-x86_64.zip 
$ unzip protoc-3.19.4-osx-x86_64.zip -d /usr/local/protobuf
$ rm protoc-3.19.4-osx-x86_64.zip 
$ echo 'export PATH="$PATH:/usr/local/protobuf/bin"' >> ~/.zshenv
```

그리고 protobuf Go 런타임 설지 
```
$ go install google.golang.org/protobuf/...@v1.27.1
```

## log.proto 

`log.proto`로 메시지를 정의한다. 
  - option go_package는 패키지명이다. 
  - 메시지 정의


## 컴파일하여 .pb.go 생성하기 

다음과 같이 컴파일 한다. 
```
$ protoc api/v1/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=.
```

1. 컴파일 할 .proto의 위치를 명시한다.  
2. `go_out`은 go로 내보내라는 말이며 컴파일한 코드 내보낼 위치를 말해준다. 
3. `go_opt`는 플래그이다. 여기서는 `paths`가 `source_relative` 하다고 한 것이다. 

```
If the paths=source_relative flag is specified, the output file is placed in the same relative directory as the input file. For example, an input file protos/buzz.proto results in an output file at protos/buzz.pb.go.
```
4. `proto_path`는 기존의 `.proto`를 가져오는 경로로 보인다. 
```
The protocol compiler searches for imported files in a set of directories specified on the protocol compiler command line using the -I/--proto_path flag. If no flag was given, it looks in the directory in which the compiler was invoked. In general you should set the --proto_path flag to the root of your project and use fully qualified names for all imports.
```

## Makefile

위 컴파일과 테스트에 대하여 미리 Makefile을 만들어두자. 

```
compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.
test:
	go test -race ./...
```
