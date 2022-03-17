package clue

import (
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CallbackReqeust 回传有效线索 API Request
type CallbackRequest struct {
	// AdvertiserIDs 广告主id列表
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// ClueID 线索ID
	ClueID string `json:"clue_id,omitempty"`
	// Source 广告来源，新版深度转化 source为1 允许值：0,1；对于source为0的线索回传，飞鱼平台不再披露显示，允许值0将于2020年9月24日下线
	Source int `json:"source,omitempty"`
	// ClueConvertState 线索状态 ：194:回访-信息确认、195:回访-加为好友、196:回访-高潜成交
	ClueConvertState int `json:"clue_convert_state,omitempty"`
}

// Encode implement PostRequest interface
func (r CallbackRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
