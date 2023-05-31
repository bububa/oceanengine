package wechat

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InstanceDetailRequest 获取微信号码包详情 API Request
type InstanceDetailRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 微信号码包ID
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// Encode implement GetRequest interface
func (r InstanceDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("instance_id", strconv.FormatUint(r.InstanceID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InstanceDetailResponse 获取微信号码包详情 API Response
type InstanceDetailResponse struct {
	model.BaseResponse
	Data *Instance `json:"instance,omitempty"`
}

// InstanceSubType 号码包类型
type InstanceSubType string

const (
	// InstanceSubType_NORMAL 普通号码包
	InstanceSubType_NORMAL InstanceSubType = "NORMAL"
	// InstanceSubType_SINGLE 单一微信号码包
	InstanceSubType_SINGLE InstanceSubType = "SINGLE"
	// InstanceSubType_ENTERPRISE_DEFAULT 企业微信
	InstanceSubType_ENTERPRISE_DEFAULT InstanceSubType = "ENTERPRISE_DEFAULT"
	// InstanceSubType_ENTERPRISE_CODE 企业微信活码
	InstanceSubType_ENTERPRISE_CODE InstanceSubType = "ENTERPRISE_CODE"
)

// Instance 微信号码包
type Instance struct {
	// InstanceID 微信号码包ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 微信号码包名称
	Name string `json:"name,omitempty"`
	// SubType 号码包类型。 可选值:
	// NORMAL: 普通号码包
	// SINGLE: 单一微信号码包
	// ENTERPRISE_DEFAULT: 企业微信
	// ENTERPRISE_CODE: 企业微信活码
	SubType InstanceSubType `json:"sub_type,omitempty"`
	// CreateTime 创建时间，格式：2022-07-21 15:36:09
	CreateTime string `json:"create_time,omitempty"`
	// ModTime 修改时间，格式：2022-07-21 15:36:09
	ModTime string `json:"mod_time,omitempty"`
	// Suffix 缀 可选值:
	// FALSE: 未开启随机尾缀
	// TRUE: 开启随机尾缀
	Suffix string `json:"suffix,omitempty"`
	// WechatList 微信号列表
	WechatList []Wechat `json:"wechat_list,omitempty"`
}
