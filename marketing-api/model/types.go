package model

import "strconv"

type AdVersion int

const (
	AdVersion_DEFAULT AdVersion = 1
	AdVersion_2       AdVersion = 2
)

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = Uint64(i)
	return
}

// Int64 support string quoted number in json
type Int64 int64

// UnmarshalJSON implement json Unmarshal interface
func (i64 *Int64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseInt(string(b), 10, 64)
	*i64 = Int64(i)
	return
}

// Float64 support string quoted number in json
type Float64 float64

// UnmarshalJSON implement json Unmarshal interface
func (f64 *Float64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseFloat(string(b), 64)
	*f64 = Float64(i)
	return
}

type OnOffInt int

func (ooi *OnOffInt) UnmarshalJSON(b []byte) (err error) {
	str := string(b)
	var i int
	if str == "ON" {
		i = 1
	} else if str == "OFF" {
		i = 0
	} else {
		if b[0] == '"' && b[len(b)-1] == '"' {
			b = b[1 : len(b)-1]
		}
		i, _ = strconv.Atoi(str)
	}
	*ooi = OnOffInt(i)
	return nil
}
