package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ExtendPackageCreateRequest 创建应用分包 API Request
type ExtendPackageCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PackageID 应用包ID，获取方法见接口文档【查询应用信息】
	PackageID string `json:"package_id,omitempty"`
	// Mode 分包模式，Auto自动生成渠道号Manual自定义渠道号信息
	Mode string `json:"mode,omitempty"`
	// ChannelCount 创建数量，（mode=Auto时需指定）单次分包取值范围1~100
	ChannelCount int `json:"channel_count,omitempty"`
	// ChannelList 自定义渠道号信息，（mode=Manual时需指定），一次调用，list的size取值范围1~100
	ChannelList []ExtendPackageChannel `json:"channel_list,omitempty"`
}

// ExtendPackageChannel 自定义渠道号信息
type ExtendPackageChannel struct {
	// ChannelID 渠道号，渠道号ID支持英文，数字，下划线和连字符-，不超过50个字符，超出部分会被截断，示例：oceanengine_ads_sample-12
	ChannelID string `json:"channel_id,omitempty"`
	// Remark 备注，渠道包备注信息，至多20个字符，超出部分会被截断处理
	Remark string `json:"remark,omitempty"`
}

// Encode implement GetRequest interface
func (r ExtendPackageCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ExtendPackageCreateResponse 创建应用分包 API Response
type ExtendPackageCreateResponse struct {
	model.BaseResponse
	Data struct {
		// PackageID 应用包id
		PackageID string `json:"package_id,omitempty"`
	} `json:"data,omitempty"`
}
