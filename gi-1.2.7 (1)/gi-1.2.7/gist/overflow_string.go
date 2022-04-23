// Code generated by "stringer -type=Overflow"; DO NOT EDIT.

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
	_ = x[OverflowAuto-0]
	_ = x[OverflowScroll-1]
	_ = x[OverflowVisible-2]
	_ = x[OverflowHidden-3]
	_ = x[OverflowN-4]
}

const _Overflow_name = "OverflowAutoOverflowScrollOverflowVisibleOverflowHiddenOverflowN"

var _Overflow_index = [...]uint8{0, 12, 26, 41, 55, 64}

func (i Overflow) String() string {
	if i < 0 || i >= Overflow(len(_Overflow_index)-1) {
		return "Overflow(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Overflow_name[_Overflow_index[i]:_Overflow_index[i+1]]
}

func (i *Overflow) FromString(s string) error {
	for j := 0; j < len(_Overflow_index)-1; j++ {
		if s == _Overflow_name[_Overflow_index[j]:_Overflow_index[j+1]] {
			*i = Overflow(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Overflow")
}
