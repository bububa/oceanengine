package enum

// AudienceGender 定向性别
type AudienceGender string

const (
	// AudienceGender_NONE 不限
	AudienceGender_NONE AudienceGender = "NONE"
	// AudienceGender_UNLIMITED 不限
	AudienceGender_UNLIMITED AudienceGender = "GENDER_UNLIMITED"
	// AudienceGender_MALE 男
	AudienceGender_MALE AudienceGender = "GENDER_MALE"
	// AudienceGender_FEMALE 女
	AudienceGender_FEMALE AudienceGender = "GENDER_FEMALE"
)
