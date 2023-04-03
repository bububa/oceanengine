package abtest

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建实验 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TestName 实验名称，最多100个字符，中文算两个字符，不支持emoji，不支持与现有实验重名。
	TestName string `json:"test_name,omitempty"`
	// TestStartTime 实验开始时间，格式为："2020-12-25 07:12:08"，开始时间不能早于当前时间。
	TestStartTime string `json:"test_start_time,omitempty"`
	// TestEndTime 实验结束时间，格式为："2020-12-25 07:12:08"，
	TestEndTime string `json:"test_end_time,omitempty"`
	// TestType 实验类型，允许值："CAMPAIGN"：广告组分流实验；"AD"：广告计划分流实验。
	TestType enum.AbTestType `json:"test_type,omitempty"`
	// ObjectList 实验对象列表
	ObjectList []AbTestObject `json:"object_list,omitempty"`
}

func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// TestID 实验ID
		TestID uint64 `json:"test_id,omitempty"`
	} `json:"data,omitempty"`
}
