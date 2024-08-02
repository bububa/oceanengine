package agent

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserSelectRequest 代理商管理账户列表 API Request
type AdvertiserSelectRequest struct {
	// AdvertiserID 代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CompanyIDs
	CompanyIDs []string `json:"company_ids,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量.默认值: 100
	PageSize int `json:"page_size,omitempty"`
	// Cursor 页码游标值，第一次拉取，无需入参
	// 注：page+page_size与cursor+count为两种分页方式
	// cursor+count适用于获取数据记录数≥10000的场景
	Cursor int `json:"cursor,omitempty"`
	// Count 页面数据量，页面数据量
	// 注：page+page_size与cursor+count为两种分页方式
	// cursor+count适用于获取数据记录数≥10000的场景
	Count int `json:"count,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserSelectRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.CompanyIDs) > 0 {
		bs, _ := json.Marshal(r.CompanyIDs)
		values.Set("company_ids", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Cursor > 0 || r.Count > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdvertiserSelectResponse 代理商管理账户列表 API Response
type AdvertiserSelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserSelectResponseData `json:"data,omitempty"`
}

// AdvertiserSelectResponseData json返回值
type AdvertiserSelectResponseData struct {
	// List 广告主ID列表
	List []uint64 `json:"list,omitempty"`
	// AccountSource 账号列表的账号类型
	// 枚举值：AD 广告主、STAR 星图、LUBAN 鲁班、DOMESTIC 入海、LOCAL 本地推账户
	AccountSource string `json:"account_source,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CursorInfo 分页信息
	CursorInfo *model.PageInfo `json:"cursor_info,omitempty"`
}
