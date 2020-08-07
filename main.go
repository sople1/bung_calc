package main

import (
	"bung_calc/tool"
	"fmt"
)

func main() {
	var _all, _passed int
	fmt.Println(":: 붕지노 포인트 계산기 ::")
	fmt.Print("\n")

	_phase := 1
	for {
		fmt.Printf("%d회차 >>>>>>>>>>>>>>\n", _phase)
		fmt.Print("=======================\n")
		_all = tool.ScanCount("총", 1, 0)
		if _all < 0 {
			tool.SetPrintPadInvert(false)
			_all = _all * -1
		} else {
			tool.SetPrintPadInvert(true)
		}
		tool.PrintPointToGet(_all)

		_passed = tool.ScanCount("통과", 1, _all)
		tool.PrintPointResult(_all, _passed)

		_phase = _phase + 1
		fmt.Print("=======================\n\n\n")
	}
}
