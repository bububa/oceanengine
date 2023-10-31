package subscribe

import "github.com/bububa/oceanengine/marketing-api/util"

// AccountsAddRequest 新增 Adv 订阅 API Request
type AccountsAddRequest struct {
	// AppID 开放平台调用MAPI/订阅RDS时所使用的APPID
	AppID uint64 `json:"app_id,omitempty"`
	// SubscribeTaskID 创建成功的RDS订阅任务所生成的订阅ID
	SubscribeTaskID uint64 `json:"subscribe_task_id,omitempty"`
	// Events 子事件扩展字段，可不填
	Events []string `json:"events,omitempty"`
	// CoreUserID 授权用户ID，获取方式https://open.oceanengine.com/labels/7/docs/1696710507039756
	CoreUserID uint64 `json:"core_user_id,omitempty"`
	// AdvertiserIDs 授权广告主账号ID
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r AccountsAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
