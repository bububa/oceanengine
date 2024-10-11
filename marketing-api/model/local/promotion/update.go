package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新广告 API Request
type UpdateRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// Name 广告名称，长度是1-50个字（两个英文字符占1个字）
	Name string `json:"name,omitempty"`
	// AwemeID 用于推广的抖音号id
	// 当营销场景为短视频，且选择素材库和上传视频投广时，该字段必填
	AwemeID string `json:"aweme_id,omitempty"`
	// VideoHpVisibility 主页视频是否可见，默认单次展示可见
	// ALWAYS_VISIBLE 主页始终可见
	// HIDE_VIDEO_ON_HP 仅单次展示可见（默认值）
	// 仅针对素材库和上传视频生效
	VideoHpVisibility enum.VideoHpVisibility `json:"video_hp_visibility,omitempty"`
	// CustomerMaterialList 视频素材列表
	// 推直播：live_material_type=VIDEO短视频时必填，live_material_type=LIVE时不支持传入，传入会报错；
	// 推短视频：未开启团购卡时，视频素材必传；开启团购卡时，可根据是否需要推广短视频素材选择传入
	CustomerMaterialList []CustomerMaterial `json:"customer_material_list,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
