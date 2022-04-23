// Code generated by "stringer -type=LineCaps"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LineCapButt-0]
	_ = x[LineCapRound-1]
	_ = x[LineCapSquare-2]
	_ = x[LineCapCubic-3]
	_ = x[LineCapQuadratic-4]
	_ = x[LineCapsN-5]
}

const _LineCaps_name = "LineCapButtLineCapRoundLineCapSquareLineCapCubicLineCapQuadraticLineCapsN"

var _LineCaps_index = [...]uint8{0, 11, 23, 36, 48, 64, 73}

func (i LineCaps) String() string {
	if i < 0 || i >= LineCaps(len(_LineCaps_index)-1) {
		return "LineCaps(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LineCaps_name[_LineCaps_index[i]:_LineCaps_index[i+1]]
}

func (i *LineCaps) FromString(s string) error {
	for j := 0; j < len(_LineCaps_index)-1; j++ {
		if s == _LineCaps_name[_LineCaps_index[j]:_LineCaps_index[j+1]] {
			*i = LineCaps(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: LineCaps")
}