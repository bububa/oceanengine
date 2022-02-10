package enum

// DmpCustomType 人群类型
type DmpCustomType int

const (
	// DmpCustomType_IMEI IMEI，大部分是15位数字，如123456789012345（电信是14位）
	DmpCustomType_IMEI DmpCustomType = 1
	// DmpCustomType_IDFA IDFA
	DmpCustomType_IDFA DmpCustomType = 2
	// DmpCustomType_UID UID
	DmpCustomType_UID DmpCustomType = 3
	// DmpCustomType_MOBILE mobile
	DmpCustomType_MOBILE DmpCustomType = 4
	// DmpCustomType_MAC MAC
	DmpCustomType_MAC DmpCustomType = 5
	// DmpCustomType_IMEI_MD5 IMEI-MD5
	DmpCustomType_IMEI_MD5 DmpCustomType = 11
	// DmpCustomType_IDFA_MD5 IDFA-MD5
	DmpCustomType_IDFA_MD5 DmpCustomType = 12
	// DmpCustomType_MOBILE_HASH MOBILE_HASH（sha256）
	DmpCustomType_MOBILE_HASH DmpCustomType = 14
	// DmpCustomType_OAID OAID
	DmpCustomType_OAID DmpCustomType = 15
	// DmpCustomType_OAID_MD5 OAID_MD5
	DmpCustomType_OAID_MD5 DmpCustomType = 16
)
