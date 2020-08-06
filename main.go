package main

import (
	"bung_calc/data"
	"bung_calc/tool"
	"fmt"
)

func main() {
	var _all, _passed int
	fmt.Println(":: 붕지노 포인트 계산기 ::")

	for {
		_all = tool.ScanCount("총", 1, 0)
		fmt.Printf("포인트: %v, %v\n", data.PointList(_all), data.PointListNegative(_all))

		_passed = tool.ScanCount("통과", 1, _all)
		fmt.Printf("총 %d / 통과 %d\n", _all, _passed)
		fmt.Printf("%v\n", tool.CalcPointList(_all, _passed))
		fmt.Printf("== 산출종료 ==\n\n")
	}
}
