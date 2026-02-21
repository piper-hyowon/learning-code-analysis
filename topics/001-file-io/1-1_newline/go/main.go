package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	// 현재 소스 파일의 디렉토리 경로 가져오기
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("런타임 정보를 가져올 수 없습니다")
		return
	}

	baseDir := filepath.Dir(currentFile)
	testdataPath := filepath.Join(baseDir, "testdata", "input.txt")

	fmt.Println("현재 파일 위치:", baseDir)
	fmt.Println("읽을 파일:", testdataPath)

	data, err := os.ReadFile(testdataPath)
	if err != nil {
		fmt.Println("파일 읽기 실패:", err)
		return
	}

	// 방식 1: bytes.Count - 빠름, 변환 없음([]byte 그대로 사용)
	count1 := bytes.Count(data, []byte("\n"))
	if len(data) > 0 && data[len(data)-1] != '\n' {
		count1++ // 마지막 줄에 개행이 없는 경우
	}

	// 방식 2: strings.Split - 각 줄을 배열로 활용 가능
	lines := strings.Split(string(data), "\n") // []byte → string 복사 발생
	count2 := len(lines)
	if lines[len(lines)-1] == "" { // 파일이 \n으로 끝나면 빈 문자열 제외
		count2--
	}

	// 방식 3: bufio.Scanner - 대용량 파일에 적합 (한 줄씩 읽음)
	count3 := countWithScanner(testdataPath)

	fmt.Printf("bytes.Count: %d줄\n", count1)
	fmt.Printf("strings.Split: %d줄\n", count2)
	fmt.Printf("bufio.Scanner: %d줄\n", count3)
}

func countWithScanner(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Scanner: 파일 열기 실패:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner: 읽기 오류:", err)
		return 0
	}

	return count
}
