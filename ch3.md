# Chapter 3. Write a Log Pacakge

로그는 추가만이 가능한 레코드의 연속이다. 

## store.go

원래는 `internal/server`에 `server` 패키지에 `log.go`가 있었다. 

이제는 `internal/log/store.go` 에 `log` 패키지를 만든다.

1. 저장을 하는 주체인 `store`를 정의하고 생성함수 `newStore()`를 정의한다. 
    - 대략 파일을 넘겨주면, 파일, 사이즈, 버퍼, 뮤텍스를 가진 `store`를 생성한다.

2. 이에 Append와 Read 메서드를 구현한다. 
    - `Append(p []byte)`에서 p는 뭐의 약자지? 그냥 b가 낫지 않나? 
    - 길이정보를 먼저 쓰고, 데이터를 쓰기에 `w += lenWidth`인거다. 

    - `Read`는 Flush부터 먼저하자
    - 사이즈 정보는 쓸때는 uint64인데 읽을때는 왜 int64일까? ReadAt 시그니처가 그렇다.

3. `ReadAt`도 구현하고, Close()도 구현한다
    - Flush를 해줘야 하기에 이렇게 감싸준다.
    
## store_test.go

이제 테스트 해봐야지
- Append, Read, ReadAt을 테스트하고
- Close()의 경우는 flush를 테스트 한다. 

