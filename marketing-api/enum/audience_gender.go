package enum

type AudienceGender string

const (
	AudienceGender_NONE   AudienceGender = "NONE"          // 不限
	AudienceGender_MALE   AudienceGender = "GENDER_MALE"   // 男
	AudienceGender_FEMALE AudienceGender = "GENDER_FEMALE" // 女
)
