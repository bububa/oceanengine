package abtest

// AbTestObject 实验对象
type AbTestObject struct {
	// ObjectID 实验对象ID，当test_type为CAMPAIGN时，传入广告组ID，当test_type为AD时，传入广告计划ID。
	ObjectID uint64 `json:"object_id,omitempty"`
}
