// Code generated by "stringer -type=DialogState"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DialogExists-0]
	_ = x[DialogOpenModal-1]
	_ = x[DialogOpenModeless-2]
	_ = x[DialogAccepted-3]
	_ = x[DialogCanceled-4]
	_ = x[DialogStateN-5]
}

const _DialogState_name = "DialogExistsDialogOpenModalDialogOpenModelessDialogAcceptedDialogCanceledDialogStateN"

var _DialogState_index = [...]uint8{0, 12, 27, 45, 59, 73, 85}

func (i DialogState) String() string {
	if i < 0 || i >= DialogState(len(_DialogState_index)-1) {
		return "DialogState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DialogState_name[_DialogState_index[i]:_DialogState_index[i+1]]
}

func (i *DialogState) FromString(s string) error {
	for j := 0; j < len(_DialogState_index)-1; j++ {
		if s == _DialogState_name[_DialogState_index[j]:_DialogState_index[j+1]] {
			*i = DialogState(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: DialogState")
}