package comment

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeBandedList 获取屏蔽用户列表 API Request
type AwemeBandedListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件
	Filtering *AwemeBandedListFiltering `json:"filtering,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// AwemeBandedListFiltering 筛选条件
type AwemeBandedListFiltering struct {
	// BandType 屏蔽类型，允许值：CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
	BandType string `json:"band_type,omitempty"`
	// AwemeID 抖音id，抖音id限制长度20，纯数字id不能以0开头。
	AwemeID string `json:"aweme_id,omitempty"`
	// NicknameKeyword 昵称关键词，关键词长度不大于40个字符，中文算2个字符。
	NicknameKeyword string `json:"nickname_keyword,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeBandedListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AwemeBandedListResponse 获取屏蔽用户列表 API Response
type AwemeBandedListResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *AwemeBandedListResponseData `json:"data,omitempty"`
}

// AwemeBandedListResponseData json 返回值
type AwemeBandedListResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 结果列表
	List []BandedAweme `json:"list,omitempty"`
}

// BandedAweme 屏蔽用户
type BandedAweme struct {
	// AwemeID 抖音id
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音昵称
	AwemeName string `json:"aweme_name,omitempty"`
	// BandedType 屏蔽类型，CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
	BandedType string `json:"banded_type,omitempty"`
	// NicknameKeyword 昵称关键词
	NicknameKeyword string `json:"nickname_keyword,omitempty"`
}
