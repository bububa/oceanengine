package model

import "strconv"

type FlexUint64 uint64

func (fu64 *FlexUint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*fu64 = FlexUint64(i)
	return
}
