package oauth

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserGetRequest 获取已授权账户 API Request
type AdvertiserGetRequest struct {
	// AppId 开发者申请的应用APP_ID
	AppId uint64 `json:"app_id,omitempty"`
	// Secret 开发者应用的私钥Secret
	Secret string `json:"secret,omitempty"`
	// AccessToken 授权access_token
	AccessToken string `json:"access_token,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("app_id", strconv.FormatUint(r.AppId, 10))
	values.Set("secret", r.Secret)
	values.Set("access_token", r.AccessToken)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdvertiserGetResponse 获取已授权账户 API Response
type AdvertiserGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserGetResponseData `json:"data,omitempty"`
}

// AdvertiserGetResponseData json返回值
type AdvertiserGetResponseData struct {
	// List 广告主列表
	List []Advertiser `json:"list,omitempty"`
}

// Advertiser 广告账户
type Advertiser struct {
	// AdvertiserID 账号id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AccountStringID 账号id（字符串型）
	// 当advertiser_role=10有效，即抖音号类型时，即为aweme_sec_uid，可用于Dou+接口调用
	AccountStringID string `json:"account_string_id,omitempty"`
	// AdvertiserName 账号名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// AdvertiserRole 账号角色，1-普通广告主，2-账号管家，3-一级代理商，4-二级代理商, 6-星图账号，10-抖音号（用于Dou+接口调用）
	AdvertiserRole uint `json:"advertiser_role,omitempty"`
	// IsValid 授权有效性，允许值：true/false；false表示对应的user在客户中心/一站式平台代理商平台变更了对此账号的权限,需要到对应平台进行调整过来；
	IsValid bool `json:"is_valid,omitempty"`
	// AccountRole 新版授权账号角色
	AccountRole enum.AccountRole `json:"account_role,omitempty"`
	// CompanyList
	CompanyList []Company `json:"company_list,omitempty"`
}

// Company
type Company struct {
	// CustomerCompanyID
	CustomerCompanyID uint64 `json:"customer_company_id,omitempty"`
	// CustomerCompanyName
	CustomerCompanyName string `json:"customer_company_name,omitempty"`
}
