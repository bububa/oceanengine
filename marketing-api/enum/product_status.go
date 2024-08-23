package enum

import "strconv"

// ProductStatus 商品状态，可选值:
type ProductStatus int

const (
	// ProductStatus_STATUS_OFFLINE 线下不可投放状态
	ProductStatus_STATUS_OFFLINE ProductStatus = 0
	// ProductStatus_STATUS_ONLINE 线上可投放状态
	ProductStatus_STATUS_ONLINE ProductStatus = 1
)

func (status *ProductStatus) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	str := string(b)
	var i int
	switch str {
	case "STATUS_OFFLINE":
		i = 0
	case "STATUS_ONLINE":
		i = 1
	default:
		i, _ = strconv.Atoi(str)
	}
	*status = ProductStatus(i)
	return nil
}

func (i ProductStatus) Value() int {
	return int(i)
}

func (i ProductStatus) String() string {
	if i == 1 {
		return "STATUS_ONLINE"
	}
	return "STATUS_OFFLINE"
}
