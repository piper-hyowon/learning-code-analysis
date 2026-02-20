package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "testdata/input.txt"
	data, err := os.ReadFile(filename)
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
	count3 := countWithScanner(filename)

	fmt.Printf("bytes.Count: %d, strings.Split: %d, bufio.Scanner: %d\n", count1, count2, count3)
}

func countWithScanner(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}
