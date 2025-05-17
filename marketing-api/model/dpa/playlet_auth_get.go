package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PlayletAuthGetRequest 查询短剧商品原片授权申请状态 API Request
type PlayletAuthGetRequest struct {
	// AdvertiserID 广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID，可通过「获取商品列表」接口获取
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductID 商品ID，可通过「获取商品列表」接口获取
	ProductID uint64 `json:"product_id,omitempty"`
}

// Encode implements GetRequest interface
func (r PlayletAuthGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
	values.Set("product_id", strconv.FormatUint(r.ProductID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PlayletAuthGetResponse 查询短剧商品原片授权申请状态 API Response
type PlayletAuthGetResponse struct {
	Data *PlayletAuthGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type PlayletAuthGetResult struct {
	// AlbumID 短剧ID
	AlbumID string `json:"album_id,omitempty"`
	// Status 申请状态 可选值:
	// ING 申请中
	// SUCCESS 申请成功
	// FAIL 申请失败
	// UNSUBMITTED 未提交申请
	Status string `json:"status,omitempty"`
	// ApplyTime 申请提交时间，时间格式为yyyy-mm-dd hh:mm:ss
	ApplyTime string `json:"apply_time,omitempty"`
	// ExpireTime 授权过期时间，时间格式为yyyy-mm-dd hh:mm:ss
	ExpireTime string `json:"expire_time,omitempty"`
	// AlbumName 短剧名称
	AlbumName string `json:"album_name,omitempty"`
}
