package dpa

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoGetRequest 获取 DPA 商品库视频模板 API Request
type VideoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *VideoGetFilter `json:"filtering,omitempty"`
	// Page 页码， 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量， 默认值: 10，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// VideoGetFilter 过滤条件
type VideoGetFilter struct {
	// ProductPlatformIDs 商品库ids，长度限制 50个
	ProductPlatformIDs []uint64 `json:"product_platform_ids,omitempty"`
	// ProductIDs 商品库商品ids，长度限制 50个
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// StartTime 视频创建的筛选起始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 视频创建的筛选结束时间
	EndTime string `json:"end_time,omitempty"`
	// ImageMode 素材类型，详见【附录-枚举值-素材类型】
	// 可选值: CREATIVE_IMAGE_MODE_VIDEO：横版视频,CREATIVE_IMAGE_MODE_VIDEO_VERTICAL：竖版视频
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// PackageID 根据视频模板id搜索，精确匹配，长度限制 50字
	PackageID string `json:"package_id,omitempty"`
	// PackageName 根据视频模板名称搜索，精确匹配 长度限制 50字
	PackageName string `json:"package_name,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoGetResponse 获取 DPA 商品库视频模板 API Response
type VideoGetResponse struct {
	model.BaseResponse
	// Data 返回值
	Data *VideoGetResponseData `json:"data,omitempty"`
}

// VideoGetResponseData 返回值
type VideoGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 商品视频模板列表
	List []Video `json:"list,omitempty"`
}
