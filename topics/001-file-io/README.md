# Topic 1: File I/O

- 코드 분석하려면 파일 시스템에서 읽어와야 함
- Go: os, io, filepath, fs (1.16+) — ~~io/ioutil~~ deprecated, ~~filepath.Walk~~ → fs.WalkDir
- C#: System.IO, File, StreamReader, Directory.GetFiles, DirectoryInfo


## 미션

### Level 1: 단일 파일 조작
- 파일 전체 읽어서 라인 수 세기
  - [x] Go
  - [ ] CSharp
- 파일 생성(json, markdown)
  - [ ] Go
  - [ ] CSharp

### Level 2: 디렉토리 탐색 + 필터링
- 단일 디렉토리에서 특정 확장자 필터링 (os.ReadDir)
  - [ ] Go
  - [ ] CSharp
- 재귀 탐색으로 확장 (fs.WalkDir)
  - [ ] Go
  - [ ] CSharp
- 제외 패턴 적용 (vendor/, node_modules/, _test.go, .Designer.cs 등)
  - [ ] Go
  - [ ] CSharp
- 파일 메타데이터 수집 (크기, 수정일)
  - [ ] Go
  - [ ] CSharp

### Level 3: 통합 — 프로젝트 스캐너
- 디렉토리 구조를 트리 형태로 출력
  - [ ] Go
  - [ ] CSharp
- 지금까지 기능들 조합해서 [프로젝트 소스 파일 수집기] 만들기
  - [ ] Go
  - [ ] CSharp
  - => 최종 프로젝트의 첫 번째 모듈?

### TODO (당장은 불필요하지만 알아두기)
- [ ] 대용량 파일 bufio.Scanner 스트리밍 읽기 vs os.ReadFile 메모리 비교
  - API 자동 생성기에서는 소스코드 파일 대상이라 대용량은 거의 없음
  - 단, 생성된 파일(.g.cs 등)이 수천 줄일 수 있으므로 인지는 해둘 것

---

# 메모

## 경로 기준점
```go
// CWD 기준 (실행 위치 의존)
// 방법 1: 암묵적
filename := "testdata/input.txt"

// 방법 2: 명시적
dir, _ := os.Getwd()
filename := filepath.Join(dir, "testdata/input.txt")
// → 둘 다 cd topics/.../go 후 실행 필요

// 소스 파일 기준 (위치 독립)
_, file, _, _ := runtime.Caller(0)
dir := filepath.Dir(file)
filename := filepath.Join(dir, "testdata/input.txt")
// → 어디서든 실행 가능
```

## 1-1_newline

### Go
- os.ReadFile() / WriteFile(): 1.16+ (이전에는 ioutil)
- os.Open() 은?
  - os.ReadFile()은 내부적으로 os.Open()을 처리
    - os.Open() -> Read() -> Close()
    - 파일 전체를 한 번에 읽을 때 간단하게 사용
  - os.Open()은 큰 파일(수백MB 이상)을 조금씩 읽거나, 세밀한 제어가 필요할 때
    - 1. 파일 일부만 읽기
    - 2. 특정 위치부터 읽기 (file.Seek(1000, 0) // 1000바이트 건너뛰기)
    - 3. 한 줄씩 스트리밍(메모리 절약)
    - 4. 읽기/쓰기 모드 지정(언제 쓰이는거지?)

### 라인 수 세기 방식 비교

| 방식 | 메모리 | 속도 | 용도 |
|---|---|---|---|
| `bytes.Count` | 파일 전체 | 가장 빠름 | 단순 카운트만 |
| `strings.Split` | 파일 전체 + 슬라이스 | 중간 | 각 줄 배열로 처리 |
| `bufio.Scanner` | 한 줄씩만 | 느림 | 대용량 파일(GB), 안전성 |

- OS별로 NewLine 설정이 다름 (Unix = \n, Windows = \r\n)
- bufio.Scanner는 \n, \r\n, \r 모두 자동 처리해서 가장 안전
- bytes.Count, strings.Split 은 UTF-8 + Unix 스타일 줄바꿈 환경에서만 작동


### C#
- 줄바꿈 자동 처리
  - Environment.NewLine (읽기 전용, 플랫폼별 기본값)
  - StreamReader.ReadLine() - \r, \n, \r\n 모두 자동 인식
  
- 줄바꿈 커스터마이징
  - Console.Out.NewLine (콘솔 출력)
  - StreamWriter.NewLine (파일 쓰기)
  - 예: writer.NewLine = "\n"; // Unix 스타일 강제

---
## 주의사항
- 텍스트 파일의 인코딩 확인 필요 (UTF-8, UTF-16, EUC-KR 등)
- 마지막 줄에 newline이 없는 경우 처리 (count+1 필요할 수 있음)
- 바이너리 파일은 라인 수 세기 의미 없음