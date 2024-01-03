package nativeanchor

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetDetailRequest 获取原生锚点详情 API Request
type GetDetailRequest struct {
	// AdvertiserID 广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AnchorIDs 锚点id列表，list长度为20，注意只允许传入同一类（anchor_type)查询详情，否则报错
	// 您可从「获取原生锚点列表」获取锚点的anchor_type
	AnchorIDs []string `json:"anchor_ids,omitempty"`
	// AnchorType 锚点类型，同时只能查询一种锚点类型，可选值:
	// APP_GAME游戏
	// APP_INTERNET_SERVICE网服
	// APP_SHOP 电商
	// PRIVATE_CHAT 咨询
	// SHOPPING_CART 购物
	// INSURANCE外跳锚点
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
}

// Encode implement GetRequest interface
func (r GetDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("anchor_ids", string(util.JSONMarshal(r.AnchorIDs)))
	values.Set("anchor_type", string(r.AnchorType))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetDetailResponse 获取原生锚点详情 API Response
type GetDetailResponse struct {
	model.BaseResponse
	Data *GetDetailResult `json:"data,omitempty"`
}

type GetDetailResult struct {
	// List 锚点详情列表
	List []NativeAnchor `json:"list,omitempty"`
	// ErrorList 查询失败的锚点及查询失败原因
	ErrorList []Error `json:"error_list,omitempty"`
}

type Error struct {
	// AnchorID 锚点id
	AnchorID string `json:"anchor_id,omitempty"`
	// Message 查询失败的原因
	Message string `json:"message,omitempty"`
	// ErrorMessage 查询失败的原因
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e Error) Error() string {
	message := e.Message
	if e.ErrorMessage != "" {
		message = e.ErrorMessage
	}
	return util.StringsJoin("锚点ID: ", e.AnchorID, ", 错误信息: ", message)
}
