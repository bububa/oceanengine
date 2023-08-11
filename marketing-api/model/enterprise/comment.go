package enterprise

import (
	"encoding/json"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
	"strconv"
)

type ListItemRequest struct {
	CcAccountId uint64  `json:"cc_account_id" required:"true"` // 纵横组织ID，纵横组织管理员或协作者授权后，通过【获取已授权账户】接口，查询到“账号角色为CUSTOMER_ADMIN-管理员授权的纵横组织、或CUSTOMER_OPERATOR-协作者授权的纵横组织”的advertiser_id，即为纵横组织ID
	EDouyinId   string  `json:"e_douyin_id" required:"true"`   // 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	StartTime   string  `json:"start_time"`                    // 开始时间，默认30天前 只支持天级的。e.g.2022-01-01 开始时间必须大于2021-01-01
	EndTime     string  `json:"end_time"`                      // 结束时间，默认当天 只支持天级的。e.g.2022-01-02 结束时间必须晚于开始时间，查询时间跨度不能大于3个月
	Filter      *Filter `json:"filter"`                        // 筛选字段
	Page        int     `json:"page"`                          // 页数，默认值：1
	PageSize    int     `json:"page_size"`                     // 页面大小，支持范围1-100 之间，默认值：20
}

func (r ListItemRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("cc_account_id", strconv.FormatUint(r.CcAccountId, 10))
	values.Set("e_douyin_id", r.EDouyinId)
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.Filter != nil {
		filtering, _ := json.Marshal(r.Filter)
		values.Set("filter", string(filtering))
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

// Filter
type Filter struct {
	ItemType string `json:"item_type"` // 视频类型 允许值：ITEM_AD广告视频、ITEM_CONTENT非广告视频(抖音视频)
}

type ListItemResponse struct {
	model.BaseResponse
	Data struct {
		ItemList []*Item `json:"item_list"`
	} `json:"data"`
}

type Item struct {
	ItemType       string `json:"item_type"`        // 视频类型 允许值：ITEM_AD广告视频、ITEM_CONTENT非广告视频(抖音视频)
	MaterialID     int    `json:"material_id"`      // 广告视频素材id
	ItemID         int64  `json:"item_id"`          // 抖音视频id
	ItemTitle      string `json:"item_title"`       // 抖音视频标题
	ItemCoverURL   string `json:"item_cover_url"`   // 抖音视频封面url
	ItemDuration   int    `json:"item_duration"`    // 抖音视频时长 单位是秒
	ItemCreateTime string `json:"item_create_time"` // 抖音视频发布时间
	CommentCount   string `json:"comment_count"`    // 评论数量
}

type GetBindListRequest struct {
	AdvertiserId uint64 `json:"advertiser_id"`
}

func (r GetBindListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserId, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

type GetBindListResponse struct {
	model.BaseResponse
	Data struct {
		List []*BindItem `json:"list"`
	} `json:"data"`
}

type BindItem struct {
	OpenID         string           `json:"open_id"`         // 抖音号open_id（含企业号open_id）
	AwemeID        string           `json:"aweme_id"`        // 抖音号id（含企业号id）
	AwemeName      string           `json:"aweme_name"`      // 抖音号名称（含企业号名称）
	AwemeAvatar    string           `json:"aweme_avatar"`    // 抖音号头像（含企业号头像）
	AuthorizeTimes []AuthorizeTimes `json:"authorize_times"` // 授权有效期
}

type AuthorizeTimes struct {
	AuthorizeStartTime string `json:"authorize_start_time"` // 授权开始时间
	AuthorizeEndTime   string `json:"authorize_end_time"`   // 授权结束时间
}
