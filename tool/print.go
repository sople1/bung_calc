package tool

import (
	"bung_calc/data"
	"fmt"
)

var padPoint int = 1000
var padPos int = 1

func SetPrintPadInvert(pos bool) {
	_pos := 1
	if !pos {
		_pos = -1
	}
	padPos = _pos
}

func PrintPointToGet(count int) {
	_p := data.PointList(count)
	_np := data.PointListNegative(count)

	if padPos < 0 {
		fmt.Println(">>> 역배당 <<<")
	}

	fmt.Println("\n** 상위 포인트")
	for i := 0; i < len(_p); i = i + 1 {
		j := len(_p) - 1 - i
		fmt.Printf(">> %d등: %d point\n", i+1, _p[j]*padPoint*padPos)
	}

	fmt.Println("\n** 하위 포인트")
	for i := len(_np) - 1; i >= 0; i = i - 1 {
		fmt.Printf(">> %d등: %d point\n", i+1, _np[i]*padPoint*padPos)
	}

	fmt.Print("\n\n")
}

func PrintPointResult(allCount int, passedCount int) {
	_pList := CalcPointList(allCount, passedCount)

	fmt.Print("\n== 산출결과 ==\n\n")
	fmt.Printf("총 %d / 통과 %d\n", allCount, passedCount)

	for i := 1; i <= allCount; i = i + 1 {
		_point := 0
		if i <= passedCount {
			_point = _pList[i-1]
		}
		fmt.Printf(">> %d등: %d point\n", i, _point*padPoint*padPos)
	}

	fmt.Print("\n== 산출종료 ==\n\n")
}
