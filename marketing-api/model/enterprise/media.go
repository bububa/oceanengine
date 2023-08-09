package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
	"strconv"
)

type ListMediaRequest struct {
	EDouyinId string `json:"e_douyin_id" required:"true"` // 企业号账户ID，获取到授权的纵横组织ID后，再通过【获取纵横组织下资产账户列表（分页）】接口，查询到e_douyin_id，即为企业号账户ID，需确保传入的企业号账户ID与纵横组织ID已建立绑定关系，且绑定关系未解除
	Page      int    `json:"page"`                        // 页数，默认值：1
	PageSize  int    `json:"page_size"`                   // 页面大小，支持范围1-100 之间，默认值：20
}

func (r ListMediaRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("e_douyin_id", r.EDouyinId)
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

type ListMediaResponse struct {
	model.BaseResponse
	Data struct {
		ItemList []*Media `json:"item_list"`
	} `json:"data"`
}

type Media struct {
	MediaID   string `json:"media_id"`   // 素材ID
	MediaType string `json:"media_type"` // 素材类型 (枚举值：IMAGE 图片素材、PDF PDF素材)
	URL       string `json:"url"`        // 素材URL
	Status    string `json:"status"`     // 素材审核状态 (枚举值：INVALID 无效/逻辑删除、VALID 有效、AUDITING 审核中、AUDIT_FAIL 审核不通过)
	URI       string `json:"uri"`        // 素材URI
}
