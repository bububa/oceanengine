package audiencepackage

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// CreateRequest 创建定向包 API Request
type CreateRequest struct {
	Audience
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 定向包名称
	Name string `json:"name,omitempty"`
	// Description 定向包描述
	Description string `json:"description,omitempty"`
	// LandingType 定向包推广类型
	LandingType string `json:"landing_type,omitempty"`
	// DeliveryRange 广告投放范围【附录：广告投放范围】
	DeliverRange string `json:"deliver_range,omitempty"`
	// HideIfExists 已安装用户，0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。
	HideIfExists int `json:"hide_if_exists,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// CreateResponse 创建定向包 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	} `json:"data,omitempty"`
}
