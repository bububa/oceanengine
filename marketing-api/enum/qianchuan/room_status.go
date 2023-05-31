package qianchuan

// RoomStatus 直播间状态
type RoomStatus string

const (
	// RoomStatus_ALL 全部
	RoomStatus_ALL RoomStatus = "ALL"
	// RoomStatus_LIVING 直播中
	RoomStatus_LIVING RoomStatus = "LIVING"
	// RoomStatus_FINISH 直播结束
	RoomStatus_FINISH RoomStatus = "FINISH"
)
