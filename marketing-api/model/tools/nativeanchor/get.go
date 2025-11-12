package nativeanchor

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取账户下原生锚点
type GetRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码，默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量，默认值: 10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤器
	Filtering *GetFilter `json:"filtering,omitempty"`
}

// GetFilter 过滤器
type GetFilter struct {
	// AnchorID 按原生锚点id 做过滤
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorName 按原生锚点名称做过滤
	AnchorName string `json:"anchor_name,omitempty"`
	// Status 锚点审核状态，允许值：锚点新建：CREATE、审核通过：AUDIT_SUCCESS、审核不通过：AUDIT_FAILD
	Status []string `json:"status,omitempty"`
	// LandingType 允许值：LINK销售线索收集、APP应用、AWEME抖音号推广
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// ExternalAction 优化目标，可通过【获取优化目标】接口获取
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// Source 锚点来源，允许值：自动生成：AUTO、手动创建：MANUAL
	Source string `json:"source,omitempty"`
	// AnchorType 锚点类型，允许值：应用下载-游戏：APP_GAME、应用下载-网服：APP_INTERNET_SERVICE、应用下载-电商：APP_SHOP、高级在线预约：ONLINE_SUBSCRIBE
	AnchorType []enum.AnchorType `json:"anchor_type,omitempty"`
	// AndroidPackageName 安卓应用包名
	AndroidPackageName string `json:"android_package_name,omitempty"`
	// IosPackageName ios应用包名
	IosPackageName string `json:"ios_package_name,omitempty"`
	// AnchorStartTime 锚点创建开始日期，格式：YYYY-MM-DD
	AnchorStartTime string `json:"anchor_start_time,omitempty"`
	// AnchorEndTime 锚点创建结束日期，格式：YYYY-MM-DD
	AnchorEndTime string `json:"anchor_end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取账户下原生锚点
type GetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData 返回数据
type GetResponseData struct {
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 广告账户下原生锚点列表
	List []NativeAnchor `json:"list,omitempty"`
}
