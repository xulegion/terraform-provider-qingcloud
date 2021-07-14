// Code generated by "stringer -type SelectionOp"; DO NOT EDIT.

package constraints

import "strconv"

const (
	_SelectionOp_name_0 = "OpUnconstrained"
	_SelectionOp_name_1 = "OpMatch"
	_SelectionOp_name_2 = "OpLessThanOpEqualOpGreaterThan"
	_SelectionOp_name_3 = "OpGreaterThanOrEqualMinorOnly"
	_SelectionOp_name_4 = "OpGreaterThanOrEqualPatchOnly"
	_SelectionOp_name_5 = "OpNotEqual"
	_SelectionOp_name_6 = "OpLessThanOrEqualOpGreaterThanOrEqual"
)

var (
	_SelectionOp_index_2 = [...]uint8{0, 10, 17, 30}
	_SelectionOp_index_6 = [...]uint8{0, 17, 37}
)

func (i SelectionOp) String() string {
	switch {
	case i == 0:
		return _SelectionOp_name_0
	case i == 42:
		return _SelectionOp_name_1
	case 60 <= i && i <= 62:
		i -= 60
		return _SelectionOp_name_2[_SelectionOp_index_2[i]:_SelectionOp_index_2[i+1]]
	case i == 94:
		return _SelectionOp_name_3
	case i == 126:
		return _SelectionOp_name_4
	case i == 8800:
		return _SelectionOp_name_5
	case 8804 <= i && i <= 8805:
		i -= 8804
		return _SelectionOp_name_6[_SelectionOp_index_6[i]:_SelectionOp_index_6[i+1]]
	default:
		return "SelectionOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}