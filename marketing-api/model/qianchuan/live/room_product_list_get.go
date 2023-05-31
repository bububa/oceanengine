package live

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RoomProductListGetRequest 获取直播间商品列表 API Request
type RoomProductListGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// Fields 需要查询的消耗指标，具体可参考返回字段
	Fields []string `json:"fields,omitempty"`
	// ExplainStatus 商品状态，默认全部，可选值:
	// ALL: 全部
	// HASEXPLAIN: 已讲解
	// UNEXPLAIN: 未讲解
	ExplainStatus qianchuan.ExplainStatus `json:"explain_status,omitempty"`
	// Page 页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认20，不超过100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r RoomProductListGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("room_id", strconv.FormatUint(r.RoomID, 10))
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.ExplainStatus != "" {
		values.Set("explain_status", string(r.ExplainStatus))
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

// RoomProductListGetResponse 获取直播间商品列表 API Response
type RoomProductListGetResponse struct {
	model.BaseResponse
	Data *RoomProductListData `json:"data,omitempty"`
}

type RoomProductListData struct {
	// List 商品列表
	List []ProductStat `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ProductStat 商品数据
type ProductStat struct {
	// ProductID 商品id
	ProductID uint64 `json:"productId,omitempty"`
	// Title 商品名称
	Title string `json:"title,omitempty"`
	// ImgURL 商品缩略图
	ImgURL string `json:"img_url,omitempty"`
	// ExplainStatus  商品状态，可选值:
	// BEINGEXPLAIN: 讲解中
	// HASEXPLAIN: 已讲解
	// UNEXPLAIN: 未讲解
	ExplainStatus qianchuan.ExplainStatus `json:"explain_status,omitempty"`
	// LiveCostPerProduct 消耗
	LiveCostPerProduct float64 `json:"live_cost_per_product,omitempty"`
	// LiveProductBindTime 上架时间
	LiveProductBindTime string `json:"live_product_bind_time,omitempty"`
	// LiveProductPrice 价格
	LiveProductPrice float64 `json:"live_product_price,omitempty"`
	// LiveProductExplainCnt 讲解次数
	LiveProductExplainCnt int64 `json:"live_product_explain_cnt,omitempty"`
	// LiveProductInventory 库存
	LiveProductInventory int64 `json:"live_product_inventory,omitempty"`
	// LiveOrderRefundAmount 广告实时退款金额
	LiveOrderRefundAmount float64 `json:"live_order_refund_amount,omitempty"`
	// LiveOrderSettleAmount7d 广告7天结算金额
	LiveOrderSettleAmount7d float64 `json:"live_order_settle_amount_7d,omitempty"`
	// LiveOrderSettleCount7d 广告7天结算订单数
	LiveOrderSettleCount7d int64 `json:"live_order_settle_count_7d,omitempty"`
	// LiveOrderSettleCountRate7d 广告7天结算订单数
	LiveOrderSettleCountRate7d float64 `json:"live_order_settle_count_rate_7d,omitempty"`
	// AdLiveOrderSettleRoiPerProduct7d 广告7天结算roi
	AdLiveOrderSettleRoiPerProduct7d float64 `json:"ad_live_order_settle_roi_per_product_7d"`
	// AdLiveOrderSettleCostPerProduct7d 广告7天结算成本
	AdLiveOrderSettleCostPerProduct7d float64 `json:"ad_live_order_settle_cost_per_product_7d"`
	// LivePayOrderGmvAlias 累积GMV
	LivePayOrderGmvAlias float64 `json:"live_pay_order_gmv_alias,omitempty"`
	// LubanLivePayOrderGmv 广告GMV
	LoubanLivePayOrderGmv float64 `json:"luban_live_pay_order_gmv,omitempty"`
	// LivePayOrderCountAlias 累计成交订单
	LivePayOrderCountAlias int64 `json:"live_pay_order_count_alias,omitempty"`
	// LubanLivePayOrderCount 广告成交订单
	LubanLivePayOrderCount int64 `json:"luban_live_pay_order_count,omitempty"`
	// LivePayOrderRoiPerProduct 直接支付ROI
	LivePayOrderRoiPerProduct float64 `json:"live_pay_order_roi_per_product,omitempty"`
	// ProductClickPayUcntRatio 点击-成交转化率
	ProductClickPayUcntRatio float64 `json:"product_click_pay_ucnt_ratio,omitempty"`
	// TotalLivePayOrderGmvEcom GPM
	TotalLivePayOrderGmvEcom float64 `json:"total_live_pay_order_gpm_ecom,omitempty"`
}
