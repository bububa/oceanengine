package enum

import "strconv"

// HideIfExists 已安装用户
// 0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。
type HideIfExists string

const (
	// HideIfExists_UNLIMITED 不限（默认值）
	HideIfExists_UNLIMITED HideIfExists = "UNLIMITED"
	// HideIfExists_FILTER 过滤，仅安卓应用推广时支持，其他情况传入不生效
	HideIfExists_FILTER HideIfExists = "FILTER"
	// HideIfExists_TARGETING 定向
	HideIfExists_TARGETING HideIfExists = "TARGETING"
)

// UnmarshalJSON implement json Unmarshal interface
func (he *HideIfExists) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	str := string(b)
	if ret, err := strconv.Atoi(str); err == nil {
		switch ret {
		case 0:
			*he = HideIfExists_UNLIMITED
		case 1:
			*he = HideIfExists_FILTER
		case 2:
			*he = HideIfExists_TARGETING
		}
	} else {
		*he = HideIfExists(str)
	}
	return
}

func (he HideIfExists) Value() int {
	switch he {
	case HideIfExists_FILTER:
		return 1
	case HideIfExists_TARGETING:
		return 2
	default:
		return 0
	}
}

func (he HideIfExists) String() string {
	if he == "" {
		return string(HideIfExists_UNLIMITED)
	}
	return string(he)
}
