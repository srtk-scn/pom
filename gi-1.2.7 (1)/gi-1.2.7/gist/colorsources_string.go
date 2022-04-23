// Code generated by "stringer -type=ColorSources"; DO NOT EDIT.

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
	_ = x[SolidColor-0]
	_ = x[LinearGradient-1]
	_ = x[RadialGradient-2]
	_ = x[ColorSourcesN-3]
}

const _ColorSources_name = "SolidColorLinearGradientRadialGradientColorSourcesN"

var _ColorSources_index = [...]uint8{0, 10, 24, 38, 51}

func (i ColorSources) String() string {
	if i < 0 || i >= ColorSources(len(_ColorSources_index)-1) {
		return "ColorSources(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ColorSources_name[_ColorSources_index[i]:_ColorSources_index[i+1]]
}

func (i *ColorSources) FromString(s string) error {
	for j := 0; j < len(_ColorSources_index)-1; j++ {
		if s == _ColorSources_name[_ColorSources_index[j]:_ColorSources_index[j+1]] {
			*i = ColorSources(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: ColorSources")
}