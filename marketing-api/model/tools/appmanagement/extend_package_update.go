package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ExtendPackageUpdateRequest 更新应用子包版本 API Request
type ExtendPackageUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用包ID，获取方法见接口文档【查询应用信息】
	PackageID string `json:"package_id,omitempty"`
	// Mode 分包模式，Auto自动生成渠道号Manual自定义渠道号信息
	Mode string `json:"mode,omitempty"`
	// ChannelList 自定义渠道号信息，（mode=Manual时需指定），一次调用，list的size取值范围1~100
	ChannelList []ExtendPackageChannel `json:"channel_list,omitempty"`
}

// Encode implement GetRequest interface
func (r ExtendPackageUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
