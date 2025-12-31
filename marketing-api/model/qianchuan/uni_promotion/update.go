package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 编辑全域推广计划 API Request
type UpdateRequest struct {
	// ProductChannelInfo 渠道品信息
	ProductChannelInfo *ProductInfo `json:"product_channel_info,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// CreativeSetting 创意设置
	CreativeSetting *CreativeSetting `json:"creative_setting,omitempty"`
	// ProgrammaticCreativeMediaList 程序化创意信息
	ProgrammaticCreativeMediaList []ProgrammaticCreativeMedia `json:"programmatic_creative_media_list,omitempty"`
	// MultiProductCreativeList 商品全域创意素材信息
	// 支持0-100个素材+0-30个标题+1个投放卡片
	// 注意：对于无号商家，仅支持投放商品卡体裁
	MultiProductCreativeList []MultiProductCreative `json:"multi_product_creative_list,omitempty"`
	// AdvertiserID 千川广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 计划名称，长度为1-100个字符，其中1个汉字算2位字符。
	// 注意：
	// 1、名称不可重复，否则会报错
	// 2、仅当marketing_goal=VIDEO_PROM_GOODS时支持
	Name string `json:"name,omitempty"`
	// AdID 全域推广计划id
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implements PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 编辑全域推广计划 API Response
type UpdateResponse struct {
	Data *UpdateResult `json:"data,omitempty"`
	model.BaseResponse
}

type UpdateResult struct {
	// ErrorList 错误信息
	ErrorList []Error `json:"error_list,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
}
