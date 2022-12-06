package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取账户下创意列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *GetFiltering `json:"filtering,omitempty"`
	// Page 当前页码: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小 默认值: 10， 取值范围：1-1000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
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

// GetFiltering 过滤条件
type GetFiltering struct {
	// AdIDs 按计划ID过滤，list长度限制 1-100
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CreativeID 按创意ID过滤
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeMaterialMode 按创意呈现方式过滤，允许值：CUSTOM_CREATIVE 自定义创意、PROGRAMMATIC_CREATIVE 程序化创意
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// AdName 按计划名称过滤，长度为1-30个字符
	AdName string `json:"ad_name,omitempty"`
	// Status 按创意状态过滤，不传入即默认返回“所有不包含已删除”，其他规则详见【附录-创意查询状态】
	Status qianchuan.CreativeStatusForSearch `json:"status,omitempty"`
	// MarketingGoal 广告组营销目标，允许值：VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignID 按广告组ID过滤
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CreativeCreateStartDate 创意创建开始时间，格式："yyyy-mm-dd"
	CreativeCreateStartDate string `json:"creative_create_start_date,omitempty"`
	// CreativeCreateEndDate 创意创建结束时间，与ad_create_start_date搭配使用，格式："yyyy-mm-dd"，时间跨度不能超过180天
	CreativeCreateEndDate string `json:"creative_create_end_date,omitempty"`
	// CreativeModifyTime 创意修改时间，精确到小时，格式："yyyy-mm-dd HH"
	CreativeModifyTime string `json:"creative_modify_time,omitempty"`
}

// GetResponse 获取账户下创意列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo model.PageInfo `json:"page_info,omitempty"`
	// List 计划列表
	List []Creative `json:"list,omitempty"`
}
