package file

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PlayableCreateRequest 上传试玩/直玩素材 API Request
type PlayableCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppMaterialID 试玩/直玩素材Material_ID，可通过抖音开放平台 - 游戏管理中获取
	AppMaterialID string `json:"app_material_id,omitempty"`
	// AppID 小游戏资产的app_id
	// 可通过【获取字节小游戏】接口获取，支持上传直玩素材需要设置关联的小游戏资产ID
	AppID string `json:"app_id,omitempty"`
	// MaterialType 素材类型，可选值:
	// 直玩素材 STRAIGHT_PLAYABLE
	// 试玩素材 PLAYABLE_NEW
	MaterialType enum.MaterialType `json:"material_type,omitempty"`
}

// Encode implements PostRequest interface
func (r PlayableCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PlayableCreateResponse 上传试玩/直玩素材 API Response
type PlayableCreateResponse struct {
	model.BaseResponse
	Data *PlayableMaterial `json:"data,omitempty"`
}

type PlayableMaterial struct {
	// MaterialID 试玩/直玩素材的素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// AppID 小游戏资产的app_id
	AppID string `json:"app_id,omitempty"`
	// Signature 试玩/直玩素材的md5值
	Signature string `json:"signature,omitempty"`
	// Version 试玩/直玩素材版本信息
	Version string `json:"version,omitempty"`
	// AppPlayURI 试玩/直玩素材uri（用于广告创编时设置试玩/直玩素材）
	AppPlayURI string `json:"app_play_uri,omitempty"`
	// AppMaterialID 试玩/直玩素材Material_ID
	AppMaterialID string `json:"app_material_id,omitempty"`
	// Title 素材名称
	Title string `json:"title,omitempty"`
	// Status 素材状态
	Status int `json:"status,omitempty"`
	// CreateTime 素材创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ImageMode 素材类型
	ImageMode string `json:"image_mode,omitempty"`
	// AuditStatus 审核状态
	AuditStatus string `json:"audit_status,omitempty"`
}
