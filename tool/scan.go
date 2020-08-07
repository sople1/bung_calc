package tool

import (
	"fmt"
)

func ScanCount(text string, min int, max int) int {
	_count := 0
	_pos := 1
	for {
		fmt.Printf("%s: ", text)
		_, _ = fmt.Scanln(&_count)

		if _count < 0 {
			_pos = -1
		}

		if _count*_pos < min {
			fmt.Printf("%s의 최소값은 %d 입니다.\n", text, min)
		} else if max != 0 && _count*_pos > max {
			fmt.Printf("%s의 최대값은 %d 입니다.\n", text, max)
		} else {
			break
		}
	}

	return _count
}
