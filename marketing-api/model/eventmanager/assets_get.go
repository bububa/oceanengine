package eventmanager

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetsGetRequest 获取已创建资产列表 API Request
type AssetsGetRequest struct {
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetType 资产类型，允许值：THIRD_EXTERNAL：三方落地页
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// Filtering 过滤条件
	Filtering *AssetsGetFiltering `json:"filtering,omitempty"`
	// SortType 排序方式，允许值：ASC：升序 DESC：降序
	// 默认值ASC
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10，最大30
	PageSize int `json:"page_size,omitempty"`
}

// AssetsGetFiltering 过滤条件
type AssetsGetFiltering struct {
	// LandingPage 三方落地页数据
	LandingPage *AssetBaseInfo `json:"landing_page,omitempty"`
	// QuickApp 快应用数据
	QuickApp *QuickApp `json:"quick_app,omitempty"`
	// App 应用数据
	App *App `json:"app,omitempty"`
	// MiniProgram 字节小程序快应用资产
	MiniProgram *AssetBaseInfo `json:"mini_program,omitempty"`
}

// Encode implement GetRequest interface
func (r AssetsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("asset_type", string(r.AssetType))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AssetsGetResponse 获取已创建资产列表 API Response
type AssetsGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AssetsGetResponseData `json:"data,omitempty"`
}

// AssetsGetResponseData json返回值
type AssetsGetResponseData struct {
	// LandingPages 三方数据集合
	LandingPages []LandingPage `json:"landing_pages,omitempty"`
	// QuickApp 快应用数据
	QuickApp []QuickApp `json:"quick_app,omitempty"`
	// App 应用数据`
	App []App `json:"app,omitempty"`
	// MiniProgram 字节小程序快应用资产
	MiniProgram []MiniProgram `json:"mini_program,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
