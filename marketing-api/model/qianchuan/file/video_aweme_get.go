package file

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoAwemeGetRequest 获取抖音号下的视频 API Request
type VideoAwemeGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 需拉取视频的抖音号
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// Filtering 视频过滤条件
	Filtering *VideoAwemeGetFiltering `json:"filtering,omitempty"`
	// Cursor 页码游标值，第一次拉取，无需入参
	Cursor int `json:"cursor,omitempty"`
	// Count 页面大小，默认值30，限制1-50
	Count int `json:"count,omitempty"`
}

// VideoAwemeGetFiltering 筛选条件
type VideoAwemeGetFiltering struct {
	//ProductID 商品ID，查询关联商品的相应视频，仅短视频带货场景需入参
	ProductID uint64 `json:"product_id,omitempty"`
	// AwemeItemURL 抖音主页视频url
	AwemeItemURL string `json:"aweme_item_url,omitempty"`
	// MaterialIDs 素材id，抖音主页视频用来投放才会有，限制0-50
	// 注意：material_ids和aweme_item_id只能选择一个进行过滤，否则可能会查询不到数据
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
	// AwemeItemIDs 抖音主页视频id，限制0-50
	// 注意：material_ids和aweme_item_id只能选择一个进行过滤，否则可能会查询不到数据
	AwemeItemIDs []uint64 `json:"aweme_item_id,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoAwemeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoAwemeGetResponse 获取抖音号下的视频 API Response
type VideoAwemeGetResponse struct {
	model.BaseResponse
	Data *VideoAwemeGetResponseData `json:"data,omitempty"`
}

// VideoAwemeGetResponseData json返回值
type VideoAwemeGetResponseData struct {
	// List 视频列表
	List []Video `json:"video_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
