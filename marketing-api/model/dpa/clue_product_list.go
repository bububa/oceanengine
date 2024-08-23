package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ClueProductListRequest 获取升级版商品列表 API Request
type ClueProductListRequest struct {
	// Filtering 过滤条件
	Filtering *ClueProductListFilter `json:"filtering,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，传参范围[1,100]
	PageSize int `json:"page_size,omitempty"`
}

type ClueProductListFilter struct {
	// ProductIDs 商品ID精确搜索
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// ProductName 商品名称模糊搜索
	ProductName string `json:"product_name,omitempty"`
	// AuditStatus 商品审核状态 可选值:
	// AUDIT_STATUS_APPROVE 审核通过
	// AUDIT_STATUS_INIT 审核中
	// AUDIT_STATUS_REJECT 审核未通过
	AuditStatus enum.ProductAuditStatus `json:"audit_status,omitempty"`
	// CategoryIDs 类目id，会级联查询所有叶子类目
	CategoryIDs []uint64 `json:"category_ids,omitempty"`
	// CategoryName 类目名称，支持模糊搜索
	CategoryName string `json:"category_name,omitempty"`
	// CompletionStatus 商品必填字段完整性
	CompletionStatus []enum.ProductCompletionStatus `json:"completion_status,omitempty"`
	// ProductIDOrNameSearch 商品ID或商品名称查询
	ProductIDOrNameSearch string `json:"product_id_or_name_search,omitempty"`
	// Statuses 可投状态过滤
	Statuses []string `json:"statuses,omitempty"`
	// Rels 商品权限关系过滤，允许值：
	// REL_COP 授权
	// REL_OWN own
	Rels []string `json:"rels,omitempty"`
}

// Encode implements GetRequest interface
func (r ClueProductListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
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

// ClueProductListResponse 获取升级版商品列表 API Response
type ClueProductListResponse struct {
	Data *ClueProductListResult `json:"data,omitempty"`
	model.BaseResponse
}

type ClueProductListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Products 商品列表
	Products []Product `json:"products,omitempty"`
}
