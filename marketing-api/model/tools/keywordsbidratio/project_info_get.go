package keywordsbidratio

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ProjectInfoGetRequest 查询优词绑定的项目信息 API Request
type ProjectInfoGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PromotionWordID 优词计划ID，从【查询优词提量系数信息】接口获取
	PromotionWordID string `json:"promotion_word_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ProjectInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("promotion_word_id", r.PromotionWordID)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ProjectInfoGetResponse 查询优词绑定的项目信息 API Response
type ProjectInfoGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data struct {
		// ProjectInfo 生效的项目信息
		ProjectInfo *ProjectInfo `json:"project_info,omitempty"`
	} `json:"data,omitempty"`
}

// ProjectInfo 生效的项目信息
type ProjectInfo struct {
	// ProjectID 生效的项目Id
	ProjectID uint64 `json:"project_id,omitempty"`
	// ProjectName 生效的项目名称
	ProjectName string `json:"project_name,omitempty"`
}
