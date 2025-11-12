package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GrayGetRequest 查询白名单能力 API Request
type GrayGetRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// GrayKeys 白名单能力key，目前仅支持单次查询1个白名单能力key
	GrayKeys []enum.GrayKey `json:"gray_keys,omitempty"`
	// Version 白名单能力key所属版本，可选值:
	// OLD_VERSION 巨量广告迁移白名单（默认值）
	// UPDATED_VERSION 巨量广告升级版能力白名单，如需查询UBA、电商多直达链接等白名单能力，请传入此枚举
	Version string `json:"version,omitempty"`
	// 抖音号id
	// gray_keys = comm_roi 时，有效且必填
	AwemeIDs []uint64 `json:"aweme_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r GrayGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("gray_keys", string(util.JSONMarshal(r.GrayKeys)))
	if r.Version != "" {
		values.Set("version", r.Version)
	}
	if len(r.AwemeIDs) > 0 {
		values.Set("aweme_ids", string(util.JSONMarshal(r.AwemeIDs)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GrayGetResponse 查询白名单能力 API Response
type GrayGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 命中白名单请求
		List []GrayItem `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// GrayItem 白名单
type GrayItem struct {
	// GrayKey 白名单能力的唯一key
	GrayKey string `json:"gray_key,omitempty"`
	// InGray 当前白名单能力是否已全量，true、false。对于已经全量的key，请不要再次查询
	InGray bool `json:"in_gray,omitempty"`
	// InWhitelist 是否命中白名单。0代表命中，1代表未命中
	InWhitelist model.Bool `json:"in_whitelist,omitempty"`
	// AwemeIDs 命中白名单的抖音号列表
	AwemeIDs []uint64 `json:"aweme_ids,omitempty"`
}
