package comment

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AwemeBannedList 获取屏蔽用户列表 API Request
type AwemeBannedListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 为绑定的抖音号添加屏蔽词, 只允许传入1个，通过【获取绑定的抖音号】 接口获取，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeID string `json:"aweme_id,omitempty"`
	// IsApplyToAdv 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv bool `json:"is_apply_to_adv,omitempty"`
	// BannedType 屏蔽类型，; 允许值：CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
	BannedType string `json:"banned_type,omitempty"`
	// AwemeUserID 抖音用户id
	AwemeUserID string `json:"aweme_user_id,omitempty"`
	// NicknameKeyword 昵称关键词，最大不超过20字
	NicknameKeyword string `json:"nickname_keyword,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeBannedListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.AwemeID != "" {
		values.Set("aweme_id", r.AwemeID)
	}
	if r.IsApplyToAdv {
		values.Set("is_apply_to_adv", "true")
	}
	if r.BannedType != "" {
		values.Set("banned_type", r.BannedType)
	}
	if r.AwemeUserID != "" {
		values.Set("aweme_user_id", r.AwemeUserID)
	}
	if r.NicknameKeyword != "" {
		values.Set("nickname_keyword", r.NicknameKeyword)
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

// AwemeBannedListResponse 获取屏蔽用户列表 API Response
type AwemeBannedListResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *AwemeBannedListResponseData `json:"data,omitempty"`
}

// AwemeBannedListResponseData json 返回值
type AwemeBannedListResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 结果列表
	List []BannedAweme `json:"list,omitempty"`
}

// BannedAweme 屏蔽用户
type BannedAweme struct {
	// AwemeUserID 抖音id
	AwemeUserID string `json:"aweme_user_id,omitempty"`
	// AwemeName 抖音昵称
	AwemeName string `json:"aweme_name,omitempty"`
	// BannedType 屏蔽类型，CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
	BannedType string `json:"banned_type,omitempty"`
	// NicknameKeyword 昵称关键词
	NicknameKeyword string `json:"nickname_keyword,omitempty"`
}
