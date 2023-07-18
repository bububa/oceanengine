package qianchuan

// QcpxMode 是否开启智能优惠券
type QcpxMode string

const (
	// QCPX_MODE_ON 启用
	QCPX_MODE_ON QcpxMode = "QCPX_MODE_ON"
	// QCPX_MODE_OFF 不启用
	QCPX_MODE_OFF QcpxMode = "QCPX_MODE_OFF"
)
