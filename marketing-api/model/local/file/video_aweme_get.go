package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoAwemeGetRequest 获取抖音主页视频 API Request
type VideoAwemeGetRequest struct {
	// LocalAccountID 本地推账户id
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Filtering 过滤字段
	Filtering *VideoAwemeGetFilter `json:"filtering,omitempty"`
	// OrderField 排序字段，允许值：
	// ARRIVE_SHOP 到店量排序
	// ESTIMATE 预估效果排序
	// LIKE_CNT 点赞量排序
	// PAY_ORDER_CNT 成交量排序
	// PUBLISH_TIME 发布时间排序
	OrderField string `json:"order_field,omitempty"`
	// ExternalAction 转化目标，允许值：
	// OTO_PAY 团购购买（默认值）
	// POI_RECOMMEND 门店引流
	// 仅order_filed= ESTIMATE 预估效果排序时有效
	ExternalAction local.ExternalAction `json:"external_action,omitempty"`
	// Count 页面数据量，默认10，最大值100，最小值1
	Count int `json:"count,omitempty"`
	// Cursor 页码游标值，第一次传0，之后每次传上一次请求返回的游标值
	Cursor string `json:"cursor,omitempty"`
}

type VideoAwemeGetFilter struct {
	// AnchorInfo 视频挂载的锚点信息
	AnchorInfo *AnchorInfo `json:"anchor_info,omitempty"`
	// AwemeIDs 抖音号ids筛选，当anchor_types=ALL_ANCHOR时必传
	AwemeIDs []string `json:"aweme_ids,omitempty"`
	// ItemIDs 主页视频ids筛选，一次最大长度限制10
	ItemIDs []uint64 `json:"item_ids,omitempty"`
	// ItemStatus 素材状态筛选，默认可用 可选值:
	// ALL 全部状态(可用&不可用)
	// VALID 可用状态
	// 默认值：VALID
	ItemStatus string `json:"item_status,omitempty"`
	// StartTime 根据视频发布时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 根据视频发布时间进行过滤的结束时间，与start_time搭配使用，格式：yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty"`
}

// AnchorInfo 视频挂载的锚点信息
type AnchorInfo struct {
	// AnchorType 根据视频所挂载的锚点类型筛选（注意：请按照项目设置和下方说明传入对应的锚点类型枚举），允许值：
	// ALL_ANCHOR 不限制锚点
	// POI_ANCHOR 门店锚点
	// PRODUCT_ANCHOR 商品锚点
	// 本地推创编选择抖音主页视频范围说明：
	// 营销场景=短视频/图文，且推广目的=团购成交/门店引流/内容加热时：
	// 推广门店时，支持选择POI_ANCHOR门店锚点和PRODUCT_ANCHOR商品锚点的抖音主页视频；需同时传入anchor_types和poi_ids/product_ids，其中，门店锚点需为项目所推广的门店，商品锚点需为所推广门店下挂载的商品，门店下挂载的商品id可通过【根据门店ID查询门店下商品ID】接口获取。
	// 推广商品时，支持选择PRODUCT_ANCHOR商品锚点的抖音主页视频；需同时传入anchor_types和product_ids，其中商品锚点需为项目所推广的商品。
	// 营销场景=直播时，支持选择抖音主页下全部视频，即ALL_ANCHOR不限制锚点
	AnchorType []string `json:"anchor_type,omitempty"`
	// PoiIDs 推广的门店ids，anchor_types=POI_ANCHOR时必传
	PoiIDs []uint64 `json:"poi_ids,omitempty"`
	// ProductIDs 推广的商品ids，anchor_types=PRODUCT_ANCHOR时必传
	ProductIDs []uint64 `json:"product_ids,omitempty"`
}

func (r VideoAwemeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.ExternalAction != "" {
		values.Set("external_action", string(r.ExternalAction))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	if r.Cursor != "" {
		values.Set("cursor", r.Cursor)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoAwemeGetResponse 获取抖音主页视频 API Response
type VideoAwemeGetResponse struct {
	model.BaseResponse
	Data *VideoAwemeGetResult `json:"data,omitempty"`
}

type VideoAwemeGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.CursorInfo `json:"page_info,omitempty"`
	// VideoList 视频列表
	VideoList []VideoAweme `json:"video_list,omitempty"`
}
