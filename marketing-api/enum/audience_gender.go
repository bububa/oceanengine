package enum

// AudienceGender 定向性别
type AudienceGender string

const (
	// AudienceGender_NONE 不限
	AudienceGender_NONE AudienceGender = "NONE"
	// AudienceGender_MALE 男
	AudienceGender_MALE AudienceGender = "GENDER_MALE"
	// AudienceGender_FEMALE 女
	AudienceGender_FEMALE AudienceGender = "GENDER_FEMALE"
)
