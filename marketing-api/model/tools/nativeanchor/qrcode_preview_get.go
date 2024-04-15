package nativeanchor

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QrcodePreviewGetRequest 批量获取锚点预览url API Request
type QrcodePreviewGetRequest struct {
	// AdvertiserID 广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AnchorIDs 锚点ID列表
	AnchorIDs []string `json:"anchor_ids,omitempty"`
	// AnchorType 锚点类型
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
}

func (r QrcodePreviewGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("anchor_type", string(r.AnchorType))
	values.Set("anchor_ids", string(util.JSONMarshal(r.AnchorIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QrcodePreviewGetResponse 批量获取锚点预览url API Response
type QrcodePreviewGetResponse struct {
	model.BaseResponse
	Data *QrcodePreviewGetResult `json:"data,omitempty"`
}

type QrcodePreviewGetResult struct {
	// SuccessList 查询成功的锚点及预览url列表
	SuccessList []QrcodePreviewInfo `json:"success_list,omitempty"`
	// ErrorList 查询失败的锚点及查询失败原因
	ErrorList []Error `json:"error_list,omitempty"`
}

type QrcodePreviewInfo struct {
	// AnchorID 锚点id
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorQrcodeURL 锚点预览的二维码信息url，有效期一天（需要使用抖音App扫码），只有当锚点关联广告时，才可查询到预览url
	AnchorQrcodeURL string `json:"anchor_qrcode_url,omitempty"`
}
