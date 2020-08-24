package data

/**
10명 미만 3 (천)
15명 미만 5 3
20명 미만 7 5 3
25명 미만 10 7 5 3 1
30명 미만 15 10 7 5 3 1
*/

var pointList = []int{
	1, 3, 5, 7, 10, 15, 20,
}

func PointList(count int) []int {
	if count < 1 {
		return []int{}
	}

	_points := pointList
	_slice := 0

	switch true {
	case count < 10:
		_slice = 2
		break
	case count < 15:
		_slice = 3
		break
	case count < 20:
		_slice = 4
		break
	case count < 25:
		_slice = 5
		break
	case count < 30:
		_slice = 6
		break
	default:
		_slice = len(_points) - 1
		break
	}

	_points = _points[:_slice]

	if count < 20 {
		_points = _points[1:]
	}

	return _points
}

func PointListNegative(count int) []int {
	_points := PointList(count)
	_negativePoints := make([]int, len(_points), cap(_points))

	j := 0
	for i := len(_points) - 1; i >= 0; i = i - 1 {
		_negativePoints[j] = -_points[i]
		j = j + 1
	}

	return _negativePoints
}
