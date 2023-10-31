package enum

// LiveRoomStatus 直播间状态
type LiveRoomStatus string

const (
	// LiveRoomStatus_FINISH 结束
	LiveRoomStatus_FINISH LiveRoomStatus = "FINISH"
	// LiveRoomStatus_LIVING 直播中
	LiveRoomStatusLIVING LiveRoomStatus = "LIVING"
	// LiveRoomStatus_PAUSE 暂停
	LiveRoomStatus_PAUSE LiveRoomStatus = "PAUSE"
	// LiveRoomStatus_END 结束
	LiveRoomStatus_END LiveRoomStatus = "END"
	// LiveRoomStatus_PREPARE 未开播
	LiveRoomStatus_PREPARE LiveRoomStatus = "PREPARE"
	// LiveRoomStatus_UNKNOW 未知
	LiveRoomStatus_UNKNOW LiveRoomStatus = "UNKNOW"
)
