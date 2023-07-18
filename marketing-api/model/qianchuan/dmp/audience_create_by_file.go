package dmp

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudienceCreateByFileRequest 上传人群 API Request
type AudienceCreateByFileRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AudienceName 人群名称，限制长度30个字符
	AudienceName string `json:"audience_name,omitempty"`
	// AudienceGroup 人群分组
	// 注意：若传入的人群分组不存在，系统会自动创建
	AudienceGroup string `json:"audience_group,omitempty"`
	// DataType 文件数据类型，允许值：
	// 1: 'IMEI原值'
	// 2: 'IDFA原值'
	// 3: 'UID原值 '
	// 4: '手机号-原值'
	// 5: 'MAC地址'
	// 11: 'IMEI-MD5'
	// 12: 'IDFA-MD5'
	// 14: '手机号-SHA256'
	// 15: 'OAID'
	// 16: 'OAID-MD5'
	// 17: '手机号-MD5'
	// 18: '字节小程序OPENID'
	DataType int `json:"data_type,omitempty"`
	// MatchType 匹配方式，允许值UID：设备号匹配，定向/排除时直接按照设备号圈选，数据相对量少而精确
	MatchType string `json:"match_type,omitempty"`
	// FileKey 文件唯一标识
	FileKey string `json:"file_key,omitempty"`
}

// Encode implement PostRequest interface
func (r AudienceCreateByFileRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AudienceCreateByFileResponse 上传人群 API Response
type AudienceCreateByFileResponse struct {
	model.BaseResponse
	Data struct {
		// AudienceID 生成的人群ID
		AudienceID uint64 `json:"audience_id,omitempty"`
	} `json:"data,omitempty"`
}
