package sip

// Message 正文数据
type Message struct {
	// MessageID 唯一标识一条推送消息/数据; 64字符
	MessageID string `json:"message_id,omitempty"`
	// SubscribeTaskID 订阅任务id，订阅任务的主键; 64字符
	SubscribeTaskID string `json:"subscribe_task_id,omitempty"`
	// AdvertiserIDs 消息对应的广告主账号（全量，包含account_relation中所有map的value中的adv_id值）; min=1，max=1k
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// AccountRelation 推送广告主账号的订阅来源，是订阅的哪个账号下的广告主数据; min=1，max=1k
	AccountRelation string `json:"account_relation,omitempty"`
	// ServiceLabel 订阅服务类型
	ServiceLabel string `json:"service_label,omitempty"`
	// Data 推送数据信息，具体结构参考子文档
	Data string `json:"data,omitempty"`
	// PublishTime 消息实际产生时间
	PublishTime int64 `json:"publish_time,omitempty"`
	// Timestamp 毫秒时间戳，本条消息实际推送时间
	Timestamp int64 `json:"timestamp,omitempty"`
	// Nonce 随机数，和timestamp组合防重放
	Nonce int64 `json:"nonce,omitempty"`
}

type AccountRelation struct {
	// CoreUserIDs 广告主账号订阅来源-用户授权账号，key为订阅的用户授权账号，value为本次推送广告主列表
	CoreUserIDs map[string][]uint64 `json:"core_user_ids,omitempty"`
	// AdvertiserIDs 广告主账号订阅来源-直接订阅广告主id列表
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// CcIDs 广告主账号订阅来源-订阅管家账号id; key为订阅的管家账号，value为本次推送广告主列表
	Ccids map[string][]uint64 `json:"cc_ids,omitempty"`
}
