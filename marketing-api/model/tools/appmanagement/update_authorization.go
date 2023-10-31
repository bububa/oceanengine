package appmanagement

import "github.com/bububa/oceanengine/marketing-api/util"

// UpdateAuthorizationRequest 更新应用共享关系 API Request
type UpdateAuthorizationRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdvertiserIDs 共享关系变更的广告主对象id，不允许为空，且数量不允许大于20个。
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// BasicPackageID 应用资产id
	BasicPackageID string `json:"basic_package_id,omitempty"`
	// OperationType 共享关系变更类型。枚举值：ADD DEL 可选值:
	// ADD: 增加
	// DEL: 删除
	OperationType string `json:"operation_type,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateAuthorizationRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
