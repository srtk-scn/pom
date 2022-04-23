// Code generated by "stringer -type=BoxSides"; DO NOT EDIT.

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
	_ = x[BoxTop-0]
	_ = x[BoxRight-1]
	_ = x[BoxBottom-2]
	_ = x[BoxLeft-3]
	_ = x[BoxN-4]
}

const _BoxSides_name = "BoxTopBoxRightBoxBottomBoxLeftBoxN"

var _BoxSides_index = [...]uint8{0, 6, 14, 23, 30, 34}

func (i BoxSides) String() string {
	if i < 0 || i >= BoxSides(len(_BoxSides_index)-1) {
		return "BoxSides(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BoxSides_name[_BoxSides_index[i]:_BoxSides_index[i+1]]
}

func (i *BoxSides) FromString(s string) error {
	for j := 0; j < len(_BoxSides_index)-1; j++ {
		if s == _BoxSides_name[_BoxSides_index[j]:_BoxSides_index[j+1]] {
			*i = BoxSides(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: BoxSides")
}
