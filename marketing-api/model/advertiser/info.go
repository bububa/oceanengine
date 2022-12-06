package advertiser

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InfoRequest 广告主信息 API Request
type InfoRequest struct {
	// AdvertiserIDs 广告主ID集合(如果包含没有访问权限的ID,将返回no permission error) 取值范围: 1-100
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// Fields 查询字段集合, 默认:查询所有。字段详见下方response字段定义 .允许值: "id", "name","description", "email", "contacter", "phonenumber", "role", "status", "telephone", "address", "reason", "license_url", "license_no", "license_province", "license_city", "company", "brand", "promotion_area", "promotion_center_province", "promotion_center_city", "industry", "balance"， "create_time"
	Fields []string `json:"fields,omitempty"`
}

// Encode implement GetRequest interface
func (r InfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	fieldsBytes, _ := json.Marshal(r.Fields)
	values := util.GetUrlValues()
	values.Set("advertiser_ids", string(idsBytes))
	values.Set("fields", string(fieldsBytes))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InfoResponse 广告主信息 API Response
type InfoResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []Info `json:"data,omitempty"`
}

// Info 广告主信息
type Info struct {
	// ID 广告主ID
	ID uint64 `json:"id,omitempty"`
	// Name 账户名
	Name string `json:"name,omitempty"`
	// Description 品牌描述,即推广内容
	Description string `json:"description,omitempty"`
	// Email 联系邮箱
	Email string `json:"email,omitempty"`
	// Contacter 联系人
	Contacter string `json:"contacter,omitempty"`
	// Phonenumber 手机号码
	Phonenumber string `json:"phonenumber,omitempty"`
	// Address 固定电话
	Address string `json:"address,omitempty"`
	// LiscenseUrl 执照预览地址(链接默认1小时内有效)
	LicenseUrl string `json:"license_url,omitempty"`
	// LicenseNo 执照编号
	LicenseNo string `json:"license_no,omitempty"`
	// LicenseProvice 执照省份
	LicenseProvince string `json:"license_province,omitempty"`
	// LicenseCity 执照城市
	LicenseCity string `json:"license_city,omitempty"`
	// Company 公司名
	Company string `json:"company,omitempty"`
	// Brand 经营类别
	Brand string `json:"brand,omitempty"`
	// Role 角色
	Role enum.AdvertiserRole `json:"role,omitempty"`
	// Status 状态
	Status enum.AdvertiserStatus `json:"status,omitempty"`
	// PromotionArea 运营区域
	PromotionArea string `json:"promotion_area,omitempty"`
	// PromotionCenterProvince 运营省份
	PromotionCenterProvince string `json:"promotion_center_province,omitempty"`
	// PromotionCenterCity 运营城市
	PromotionCenterCity string `json:"promotion_center_city,omitempty"`
	// FirstIndustryName 一级行业名称（新版）
	FirstIndustryName string `json:"first_industry_name,omitempty"`
	// SecondIndustryName 二级行业名称（新版）
	SecondIndustryName string `json:"second_industry_name,omitempty"`
	// Reason 审核拒绝原因
	Reason string `json:"reason,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
}
