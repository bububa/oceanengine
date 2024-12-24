package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ClueProductDeleteRequest 删除升级版商品 API Request
type ClueProductDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProductIDs 通过巨量商品id删除，支持批量，一次请求最多允许传入100
	ProductIDs uint64 `json:"product_ids,omitempty"`
	// StoreIDAndOuterID 电商店铺商品（category_id为140000000时）支持按照店铺id+外部商品id删除，其他类目商品不需要传
	StoreIDAndOuterID *StoreIDAndOuterID `json:"store_id_and_outer_id,omitempty"`
}

// StoreIDAndOuterID 电商店铺商品
type StoreIDAndOuterID struct {
	// StoreID 店铺ID，传入时必须与商品外部ID同时传入
	StoreID string `json:"store_id,omitempty"`
	// OuterID 商品外部ID
	OuterID string `json:"outer_id,omitempty"`
}

// Encode implements PostRequest interface
func (r ClueProductDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ClueProductDeleteResponse 删除升级版商品 API Request
type ClueProductDeleteResponse struct {
	model.BaseResponse
	Data *ClueProductDeleteResult `json:"data,omitempty"`
}

type ClueProductDeleteResult struct {
	// ProductIDs 删除成功的商品ID集合
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// Errors 删除失败的商品ID集合及错误信息
	Errors []ClueProductDeleteError `json:"errors,omitempty"`
}

// Error 删除失败的商品ID集合及错误信息
type ClueProductDeleteError struct {
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// ErrorMessage 错误信息（主要的可能错误为商品关联了在投计划）
	ErrorMessage string `json:"error_message,omitempty"`
}

func (r ClueProductDeleteError) Error() string {
	return util.StringsJoin(r.ErrorMessage, "(", strconv.FormatUint(r.ProductID, 10), ")")
}
