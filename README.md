# Distributed services with Go

## 개요

책의 코드를 따라 치며 실습을 하려 한다. 

## 프로세스

1. `dswg` 메인 디렉토리를 하나 만들고 여기서 계속 진행
2. 진행중 코멘트 작업 및 필요하다면 마크다운 작업
3. 각 챕터의 코드를 완료하면 해당 시점 챕터 디렉토리를 만들어서 복사해 넣는다 

## start

```
cd dswg
go mod init github.com/nicewook/proglog
```

## 작업 일정

### 4/14(목): 1장
  
`cmd/server/ch1test.sh` 도 만들어두었음. `chmod +x ch1test.sh` 해두어야 함

```zsh
#!/usr/bin/zsh
curl -X POST localhost:8080 -d '{"record":{"value":"1234"}}'
curl -X POST localhost:8080 -d '{"record":{"value":"5678"}}'
curl -X POST localhost:8080 -d '{"record":{"value":"90AB"}}'
curl -X GET localhost:8080 -d '{"offset":0}'
curl -X GET localhost:8080 -d '{"offset":1}'
curl -X GET localhost:8080 -d '{"offset":2}'
```


