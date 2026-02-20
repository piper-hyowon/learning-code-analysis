# Topic 1: File I/O

- 코드 분석하려면 파일 시스템에서 읽어와야 함
- Go: os, io/ioutil, filepath.Walk
- C#: System.IO, File, StreamReader, Directory.GetFiles, DirectoryInfo

## 미션

### Level 1: 파일 읽기/쓰기
- 파일 전체 읽어서 라인 수 세기
  - [x] Go
  - [ ] CSharp
- 특정 확장자(.go, .cs) 파일만 필터링해서 목록 출력
  - [ ] Go
  - [ ] CSharp
- 파일 생성(json, markdown)

### Level 2: 디렉토리 탐색
- 특정 디렉토리 내 모든 소스코드 파일 재귀적으로 찾기
- 파일 메타데이터 수집(크기, 수정 날짜 등)
- 디렉토리 구조를 트리 형태로 출력


## 메모

### Go
- ioutil.ReadFile() / WriteFile(): 1.15 이전
- os.ReadFile() / WriteFile(): 1.16+
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

| 방식 | 메모리 | 용도 |
|---|---|---|
| `bytes.Count(data, []byte("\n"))` | 파일 전체 | 빠름, 단순 카운트 |
| `strings.Split(content, "\n")` | 파일 전체 + 슬라이스 | 각 줄을 배열로 다뤄야 할 때 |
| `bufio.Scanner` | 한 줄씩만 | 대용량 파일 (GB 단위)


