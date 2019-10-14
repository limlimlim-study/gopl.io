package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1은 표준 입력에서 두번 이상 나타나는 각 줄을
// 앞에 카운드를 추가해 출력한다.
func main() {
	counts := make(map[string]int)
	cnt := 0
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		cnt++
		if cnt > 10 {
			break
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
