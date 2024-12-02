package enum

import "strconv"

type OnOff string

const (
	// ON 开启
	ON OnOff = "ON"
	// OFF 关闭
	OFF OnOff = "OFF"
)

func (ooi *OnOff) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	str := string(b)
	var v string
	switch str {
	case "ON", "OFF":
		v = str
	default:
		if str != "" {
			if i, err := strconv.Atoi(str); err == nil {
				switch i {
				case 1:
					v = "ON"
				case 0:
					v = "OFF"
				}
			}
		}
	}
	*ooi = OnOff(v)
	return nil
}
