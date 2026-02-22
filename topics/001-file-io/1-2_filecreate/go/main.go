package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

type Endpoint struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

const mdTemplate = `# {{.Method}} {{.Path}}

## {{.Summary}}

{{.Description}}`

func main() {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("런타임 정보를 가져올 수 없습니다")
		return
	}

	baseDir := filepath.Dir(currentFile)
	targetDir := filepath.Join(baseDir, "testdata/output")

	data := Endpoint{
		Method:      "GET",
		Path:        "/v1/users",
		Summary:     "사용자 목록 조회",
		Description: "전체 사용자 목록 반환",
	}

	// JSON 생성
	jsonFile := filepath.Join(targetDir, "api.json")
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(jsonFile, jsonBytes, 0o644); err != nil {
		log.Fatal(err)
	}

	// Markdown 생성
	mdFile := filepath.Join(targetDir, "api.md")

	// 방법1: + 연산
	md1 := "# " + data.Method + " " + data.Path + "\n\n" +
		"## " + data.Summary + "\n\n" +
		data.Description

	// 방법2: fmt.Sprintf
	md2 := fmt.Sprintf("# %s %s\n\n## %s\n\n%s",
		data.Method, data.Path, data.Summary, data.Description)

	// 방법3: text/template
	tmpl, err := template.New("md").Parse(mdTemplate)
	if err != nil {
		log.Fatal(err)
	}
	var tmplBuf strings.Builder
	if err := tmpl.Execute(&tmplBuf, data); err != nil {
		log.Fatal(err)
	}
	md3 := tmplBuf.String()

	// 방법4: strings.Builder
	var sb strings.Builder
	sb.WriteString("# " + data.Method + " " + data.Path + "\n\n")
	sb.WriteString("## " + data.Summary + "\n\n")
	sb.WriteString(data.Description)
	md4 := sb.String()

	// 방법5: strings.Join
	md5 := strings.Join([]string{
		"# " + data.Method + " " + data.Path,
		"## " + data.Summary,
		data.Description,
	}, "\n\n")

	// 한 파일에 구분 주석 넣고 이어서 쓰기
	allContent := strings.Join([]string{
		"<!-- 방법1: + 연산 -->",
		md1,
		"<!-- 방법2: fmt.Sprintf -->",
		md2,
		"<!-- 방법3: text/template -->",
		md3,
		"<!-- 방법4: strings.Builder -->",
		md4,
		"<!-- 방법5: strings.Join -->",
		md5,
	}, "\n\n---\n\n")

	if err := os.WriteFile(mdFile, []byte(allContent), 0o644); err != nil {
		log.Fatal(err)
	}

	// 이어쓰기
	file, err := os.OpenFile(mdFile, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("\n\n---\n\n<!-- 이어쓰기: os.OpenFile -->\n> O_APPEND로 추가\n")
}
