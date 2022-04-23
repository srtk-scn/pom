// Code generated by "stringer -type=TreeViewFlags"; DO NOT EDIT.

package giv

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TreeViewFlagClosed-24]
	_ = x[TreeViewFlagChanged-25]
	_ = x[TreeViewFlagNoTemplate-26]
	_ = x[TreeViewFlagUpdtRoot-27]
	_ = x[TreeViewFlagsN-28]
}

const _TreeViewFlags_name = "TreeViewFlagClosedTreeViewFlagChangedTreeViewFlagNoTemplateTreeViewFlagUpdtRootTreeViewFlagsN"

var _TreeViewFlags_index = [...]uint8{0, 18, 37, 59, 79, 93}

func (i TreeViewFlags) String() string {
	i -= 24
	if i < 0 || i >= TreeViewFlags(len(_TreeViewFlags_index)-1) {
		return "TreeViewFlags(" + strconv.FormatInt(int64(i+24), 10) + ")"
	}
	return _TreeViewFlags_name[_TreeViewFlags_index[i]:_TreeViewFlags_index[i+1]]
}

func StringToTreeViewFlags(s string) (TreeViewFlags, error) {
	for i := 0; i < len(_TreeViewFlags_index)-1; i++ {
		if s == _TreeViewFlags_name[_TreeViewFlags_index[i]:_TreeViewFlags_index[i+1]] {
			return TreeViewFlags(i + 24), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: TreeViewFlags")
}