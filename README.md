# Go 언어를 활용한 분산 서비스 개발

## GitHub repo에 대하여

Distributed services with Go의 한글 번역서인 {책제목} 에서의 패키지, 도구의 버전과 코드 내용의 변경사항에 맞춰 실습한 내용을 챕터별로 정리하였다. 
<Go 언어를 활용한 분산 서비스 개발>에 실린 코드의 내용과 미묘하게 다를 수 있으나 큰 틀에서는 책의 내용을 따른다. 

각 챕터를 실습할 때에 해당 챕터를 go.work 파일에서 코멘트 해체하고 VSCode 에서 편집하면 편할 것이다. 

## <Go 언어를 활용한 분산 서비스 개발>의 의미

책은 분산 서비스를 Go로 구현하며 분산 서비스에 필요한 기능들과 작동의 개념을 이해하는데 의미를 가진다. 

책 자체는 두껍지 않지만 책이 소개하는 내용 전체를 소화해 내는 것은 쉽지 않을 것이다. 대상 독자를 중급 이상의 개발자로 보고 있기에 기본적인 사항은 알고 있거나 충분히 찾아서 할 수 있을 것이라 가정하고 있으며, 사용하는 패키지와 도구들인 protobuf, gRPC, raft, surf, docker, kind, helm등만 하여도 충분히 이해하려면 별도의 공부가 필요하기 때문이다.

그럼에도 불구하고, 분산 서비스라는, 매력적이지만 접근하기 쉽지않은 주제에 대해 전체를 조망해볼 수 있다는 점에서 이 책은 좋은 가이드가 될 것이다.

## 챕터별 개요와 의의

{추후 업데이트}

## 기타 안내사항

- install.md: 챕터 별 필요한 도구와 패키지들의 설치 방법 정리 
- 번역서 구매 링크: {추후 업데이트}
- 원서 웹페이지 링크: https://pragprog.com/titles/tjgo/distributed-services-with-go/
  - 책에 대한 정보와 소스코드를 얻을 수 있다. 
- 원서 코드의 오류를 바로잡는데 도움이 되었던 링크들
  - 오류 수정 게시물: https://forum.devtalk.com/t/distributed-services-with-go-unable-to-pass-readiness-liveness-checks-page-210-215/22354/3
  - 게시물의 수정 코드: https://github.com/varunbpatil/proglog/commit/2ecd0d7812b40583b76594eb148c1b4e60ca835b
  - 원서를 실습한 또 다른 GitHub repo: https://github.com/evdzhurov/dist-services-with-go/

## Special Thanks to 

- 한결같은 지지와 배려를 해준 아내와 아빠의 에너지를 채워준 두 아들
- 이 책의 번역을 소개해주신 번역 선배 김찬빈님
- 여유로운 가이드로 안전감있게 번역할 수 있도록 챙겨주신 제이펍 이상복 팀장님
- 원서 코드의 핵심 오류를 잡아내주신 권경모님
- 다양한 생각과 경험을 나누며 언제나 영감의 원천이 되는 커뮤니티들인 슬랙 <딥백수>와 디스코드 <코딩냄비> 여러분들!
