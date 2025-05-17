package enum

// ArchiveStatus 回收状态
type ArchiveStatus string

const (
	// ArchiveStatus_NO_RETRIEVE 未回收
	ArchiveStatus_NO_RETRIEVE ArchiveStatus = "NO_RETRIEVE"
	// ArchiveStatus_RETRIEVED 已回收
	ArchiveStatus_RETRIEVED ArchiveStatus = "RETRIEVED"
	// ArchiveStatus_ARCHIVED 已归档
	ArchiveStatus_ARCHIVED ArchiveStatus = "ARCHIVED"
)
