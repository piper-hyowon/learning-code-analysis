# Topic 2: Text Parsing & Regex

- 읽어온 파일 내용 분석 기초
- 정규표현식으로 코드 패턴 매칭
- Go: strings, regexp
- C#: string methods, Regex, System.Text.RegularExpressions

## 미션

### Level 1: 기본 문자열 처리
- [ ] 파일에서 특정 키워드 찾기 (func, class, public 등)
- [ ] 라인별로 읽어서 공백 라인 제거
- [ ] 대소문자 구분 없이 검색

### Level 2: 정규표현식 기초
- [ ] import/using 문 추출하기
- [ ] 단일 라인 주석 찾기 (// or #)
- [ ] 멀티 라인 주석 찾기 (/* */ or """ """)

### Level 3: 코드 패턴 매칭
- [ ] 함수/메서드 시그니처 추출 (이름, 파라미터, 리턴타입)
- [ ] 클래스/구조체 선언 찾기
- [ ] 주석 라인 vs 코드 라인 비율 계산

### Level 4: 정규식의 한계 생각해보기
- [ ] 중첩된 괄호 매칭 시도해보기
- 다음 단계(AST) 필요성 이해