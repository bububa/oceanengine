package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// IndustryInfoListRequest 获取应用细分分类及题材标签 API Request
type IndustryInfoListRequest struct {
	// AccountID 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型，允许值:
	// AD 广告主账户、BP 巨量纵横组织账号
	AccountType enum.AccountType `json:"account_type,omitempty"`
}

// Encode implement GetRequest interface
func (r IndustryInfoListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// IndustryInfoListResponse 获取应用细分分类及题材标签 API Response
type IndustryInfoListResponse struct {
	model.BaseResponse
	Data struct {
		// Industries 应用分类及题材标签信息
		Industries []Industry `json:"industries,omitempty"`
	} `json:"data,omitempty"`
}
