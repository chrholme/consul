// Code generated by "stringer -type=QueryType"; DO NOT EDIT.

package protocol

import "strconv"

const _QueryType_name = "QtNoneQtSelectQtProcedureCall"

var _QueryType_index = [...]uint8{0, 6, 14, 29}

func (i QueryType) String() string {
	if i >= QueryType(len(_QueryType_index)-1) {
		return "QueryType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QueryType_name[_QueryType_index[i]:_QueryType_index[i+1]]
}
