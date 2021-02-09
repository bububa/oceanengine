package advertiser

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type InfoRequest struct {
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"` // 广告主ID集合(如果包含没有访问权限的ID,将返回no permission error) 取值范围: 1-100
	Fields        []string `json:"fields,omitempty"`         // 查询字段集合, 默认:查询所有。字段详见下方response字段定义 .允许值: "id", "name","description", "email", "contacter", "phonenumber", "role", "status", "telephone", "address", "reason", "license_url", "license_no", "license_province", "license_city", "company", "brand", "promotion_area", "promotion_center_province", "promotion_center_city", "industry", "balance"， "create_time"
}

func (r InfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	fieldsBytes, _ := json.Marshal(r.Fields)
	values := &url.Values{}
	values.Set("advertiser_ids", string(idsBytes))
	values.Set("fields", string(fieldsBytes))
	return values.Encode()
}

type InfoResponse struct {
	model.BaseResponse
	Data []Info `json:"data,omitempty"`
}

type Info struct {
	ID              uint64                `json:"id,omitempty"`               // 广告主ID
	Name            string                `json:"name,omitempty"`             // 账户名
	Description     string                `json:"description,omitempty"`      // 品牌描述,即推广内容
	Email           string                `json:"email,omitempty"`            // 联系邮箱
	Contacter       string                `json:"contacter,omitempty"`        // 联系人
	Phonenumber     string                `json:"phonenumber,omitempty"`      // 手机号码
	Address         string                `json:"address,omitempty"`          // 固定电话
	LicenseUrl      string                `json:"license_url,omitempty"`      // 执照预览地址(链接默认1小时内有效)
	LicenseNo       string                `json:"license_no,omitempty"`       // 执照编号
	LicenseProvince string                `json:"license_province,omitempty"` // 执照省份
	LicenseCity     string                `json:"license_city,omitempty"`     // 执照城市
	Company         string                `json:"company,omitempty"`          // 公司名
	Brand           string                `json:"brand,omitempty"`            // 经营类别
	Role            enum.AdvertiserRole   `json:"role,omitempty"`             // 角色
	Status          enum.AdvertiserStatus `json:"status,omitempty"`           // 状态
}
