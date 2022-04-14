# Chapter 1. Let's Go

코드 위주로 정리해본다.

## log.go

- `Record`는 저장 최소 단위
- `Log`는 `[]Record`를 가진다. 그리고 안전하게 뮤텍스도 가진다 
- `Log`는 `Append`, `Read` 메서드를 가진다

## http.go

- `httpServer`는 `Log`를 가지고
- `handleProduce`와 `handleConsume` 핸들러 메서드를 가진다. `Log`에 바로 접근 가능
  - 그리고 Reqeust를 받아서, Response를 보내게 구현되어 있으며
  - 실제 작업은 `Log`의 메서드들이다. `Append`, `Read`

- `NewHTTPServer`가 HTTP server의 역할을 한다. 
  - mux를 사용해서 `httpServer`의 메서드를 핸들러로 연결해준다.

## ch1test.sh

- 테스트를 위한 shell script 이다. 3개의 레코드를 추가하고, 다시 읽어본다.
  - `#!/usr/bin/zsh -v`에서 -v는 명령행도 출력하라는 것이다. 

```
#!/usr/bin/zsh -v

curl -X POST localhost:8080 -d '{"record":{"value":"1234"}}'
curl -X POST localhost:8080 -d '{"record":{"value":"5678"}}'
curl -X POST localhost:8080 -d '{"record":{"value":"90AB"}}'
curl -X GET localhost:8080 -d '{"offset":0}'
curl -X GET localhost:8080 -d '{"offset":1}'
curl -X GET localhost:8080 -d '{"offset":2}'
```
