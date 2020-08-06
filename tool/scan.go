package tool

import "fmt"

func ScanCount(text string, min int, max int) int {
	count := 0
	for {
		fmt.Printf("%s: ", text)
		_, _ = fmt.Scanln(&count)

		if count < min {
			fmt.Printf("%s의 최소값은 %d 입니다.\n", text, min)
		} else if max != 0 && count > max {
			fmt.Printf("%s의 최대값은 %d 입니다.\n", text, max)
		} else {
			break
		}
	}

	return count
}
