package tool

import (
	"bung_calc/data"
)

func CalcPointList(allCount int, passedCount int) []int {
	_p := data.PointList(allCount)
	_np := data.PointListNegative(allCount)
	_list := make([]int, passedCount, allCount)

	for i := 0; i < len(_p); i = i + 1 {
		j := len(_p) - 1 - i
		if j >= len(_list) {
			continue
		}
		_list[j] += _p[i]
	}

	for i := 0; i < len(_np); i = i + 1 {
		j := passedCount - 1 - i
		if j < 0 {
			continue
		}
		_list[j] += _np[i]
	}

	return _list
}
