// Code generated by "stringer -type=SelModes"; DO NOT EDIT.

package gi3d

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NotSelectable-0]
	_ = x[Selectable-1]
	_ = x[SelectionBox-2]
	_ = x[Manipulable-3]
	_ = x[SelModesN-4]
}

const _SelModes_name = "NotSelectableSelectableSelectionBoxManipulableSelModesN"

var _SelModes_index = [...]uint8{0, 13, 23, 35, 46, 55}

func (i SelModes) String() string {
	if i < 0 || i >= SelModes(len(_SelModes_index)-1) {
		return "SelModes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SelModes_name[_SelModes_index[i]:_SelModes_index[i+1]]
}

func (i *SelModes) FromString(s string) error {
	for j := 0; j < len(_SelModes_index)-1; j++ {
		if s == _SelModes_name[_SelModes_index[j]:_SelModes_index[j+1]] {
			*i = SelModes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: SelModes")
}
