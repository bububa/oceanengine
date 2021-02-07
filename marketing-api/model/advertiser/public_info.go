package advertiser

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type PublicInfoRequest struct {
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"` // 广告主ID集合(如果包含没有访问权限的ID,将返回no permission error) 取值范围: 1-100
}

func (r PublicInfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	values := &url.Values{}
	values.Set("advertiser_ids", string(idsBytes))
	return values.Encode()
}

type PublicInfoResponse struct {
	model.BaseResponse
	Data []PublicInfo `json:"data,omitempty"`
}

type PublicInfo struct {
	ID                 uint64 `json:"id,omitempty"`                   // 广告主ID
	Name               string `json:"name,omitempty"`                 // 账户名
	Company            string `json:"company,omitempty"`              // 公司名
	FirstIndustryName  string `json:"first_industry_name,omitempty"`  // 一级行业名
	SecondIndustryName string `json:"second_industry_name,omitempty"` // 二级行业名

}
