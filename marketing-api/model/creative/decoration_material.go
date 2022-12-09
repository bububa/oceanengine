package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// DecorationMaterial 家装卡券素材
type DecorationMaterial struct {
	// ActivityID 活动ID
	ActivityID string `json:"activity_id,omitempty"`
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
}
