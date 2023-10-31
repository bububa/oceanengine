package qianchuan

// LiveroomHeatMode 直播间加热方式
type LiveroomHeatMode string

const (
	// LiveroomHeatMode_BY_ROOM 直接加热直播间
	LiveroomHeatMode_BY_ROOM LiveroomHeatMode = "BY_ROOM"
	// LiveroomHeatMode_BY_VIDEO 选择视频加热直播间
	LiveroomHeatMode_BY_VIDEO LiveroomHeatMode = "BY_VIDEO"
)
