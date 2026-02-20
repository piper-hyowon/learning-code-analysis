# Learning Code Analysis
소스 코드 분석 기술 학습 저장소

## 목표
소스 코드 분석해서 API 문서 자동 생성하는 라이브러리 개발을 위한 기반 다지기

## 구조
```
001-file-io/
├── README.md                   # 미션 설명
├── go/
│   ├── main.go
│   ├── go.mod
│   └── testdata/
│       ├── input.txt
│       └── output/            # Write 결과 저장
├── csharp/
│   ├── Program.cs
│   ├── FileIO.csproj
│   └── TestData/ 
│       ├── input.txt
│       └── output/
```

## 실행

**Go**
```
cd topics/{topic}/go
go run main.go
```


## 로드맵

### Phase 1: 기초
- [Topic 1: File I/O](./topics/001-file-io/) - 파일 시스템 다루기
- [Topic 2: Text Parsing & Regex](./topics/002-text-parsing/) - 문자열 처리 & 패턴 매칭
- [Topic 3: AST Fundamentals](./topics/003-ast-fundamentals/) - 추상 구문 트리 기초

여기까지 한 다음에 다음 계획 다시 세워보기

- Go: 정적 분석 (default 는 net/http 기준 분석, framework별은 우선 gin만 지원)
- C#: ASP.NET Core 정적 분석
- 처리할 것: Struct/DTO 필드, 태그 추출
  - HTTP 라우트 정보 수집
  - 에러 감지
  - 응답 타입 추적



## 언어
- Go
- C# 