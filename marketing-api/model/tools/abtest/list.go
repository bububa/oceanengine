package abtest

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取实验列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filter 过滤器
	Filter *ListFilter `json:"filter,omitempty"`
	// Page 页码，默认：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认：10
	PageSize int `json:"page_size,omitempty"`
}

// ListFilter 过滤器
type ListFilter struct {
	// TestIDs 根据实验ID列表过滤，列表长度1-100
	TestIDs []uint64 `json:"test_ids,omitempty"`
	// TestName 根据实验名称过滤，支持模糊匹配，长度限制100字符，中文算2个字符。
	TestName string `json:"test_name,omitempty"`
	// Conclusions 根据实验结论过滤，允许值："NOT_START"：未开始，"OBVIOUS"：有明显结论，"INSUFFICIENT"：数据量不足；"FAILED"：不满足实验条件，"TMP_OBVIOUS"：有初步结论。
	Conclusions []enum.AbTestConclusion `json:"conclusions,omitempty"`
	// Status 根据实验状态过滤，允许值："CREATED"：排期中，"PROCESSING"：进行中，"FINISH"：结束，"FAILED"：不满足实验条件。
	Status []enum.AbTestStatus `json:"status,omitempty"`
	// CreateTimeBefore 返回创建时间在此之前的实验，格式："2020-12-25 07:12:08"
	CreateTimeBefore string `json:"create_time_before,omitempty"`
	// CreateTimeAfter 返回创建时间在此之后的实验，格式："2020-12-25 07:12:08"
	CreateTimeAfter string `json:"create_time_after,omitempty"`
	// TestTimeBefore 返回测试时间在此之前的实验，格式："2020-12-25 07:12:08"
	TestTimeBefore string `json:"test_time_before,omitempty"`
	// TestTimeAfter 返回测试时间在此之后的实验，格式："2020-12-25 07:12:08"
	TestTimeAfter string `json:"test_time_after,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filter != nil {
		bs := util.JSONMarshal(r.Filter)
		values.Set("filter", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 10 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ListResponse 获取实验列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListData `json:"data,omitempty"`
}

type ListData struct {
	// TestList 实验列表
	TestList []AbTest `json:"test_list,omitempty"`
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
