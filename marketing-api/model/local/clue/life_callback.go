package clue

import (
	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LifeCallbackRequest 本地推线索回传 API Request
type LifeCallbackRequest struct {
	// LocalAccountIDs 本地推广告主id列表
	LocalAccountIDs []uint64 `json:"local_account_ids,omitempty"`
	// ClueConvertState 线索当前事件状态，可选值:
	// ARRIVAL 顾客到店/销售人员成功上门
	// CLUE_CONFIRM 顾客表达有意向
	// CLUE_HIGH_INTENTION 定金或钩子品支付
	// CONVERSION_CLASS 正价支付
	// INVALID_EVENT 无效
	ClueConvertState local.ClueConvertState `json:"clue_convert_state,omitempty"`
	// EventData 无效事件标签；若clue_convert_state = INVALID_EVENT，则该字段必填
	EventData *LifeCallbackEvent `json:"event_data,omitempty"`
	// ClueID 线索id
	ClueID string `json:"clue_id,omitempty"`
}

type LifeCallbackEvent struct {
	// EventCode 对线索当前事件状态clue_convert_state选择扩展标签。
	//  若clue_convert_state = INVALID_EVENT，则该字段必填，可选值:
	// 行业通用
	// CALL_THREE_TIMES_NO_RESPONSE 拨打3次以上未接通
	// CONNECTED_NO_INTENTION 接通但无意向
	// NO_AD_RESPONSE 反馈未在广告留资
	// OFFENSIVE_LANGUAGE 黑脏词
	// OTHER 其它
	// INVALID_NUMBER 空号/停机/关机
	// 家居行业
	// HOME_COLLABORATION_WITH_PEERS (家居)同行合作
	// HOME_NO_SHOW_FOR_MEASUREMENT_VISIT (家居)爽约上门量房
	// HOME_RENOVATION_TYPE_NOT_ACCEPTABLE (家居)装修类型不可承接
	// HOME_MISMATCH_DIFFERENT_LOCATIONS (家居)需求不匹配-需求与供给异地，原MISMATCH_DIFFERENT_LOCATIONS枚举逐步下线中
	// HOME_MISMATCH_LOW_BUDGET (家居)需求不匹配-消费力差（指代预算少/面积小），原MISMATCH_LOW_BUDGET枚举逐步下线中
	// 教育行业
	// EDU_MISMATCH_CLASS_TIME (教育)需求不匹配-上课时间不匹配
	// EDU_MISMATCH_COURSE (教育)需求不匹配-课程不匹配
	// EDU_MISMATCH_DIFFERENT_LOCATIONS (教育)需求不匹配-需求与供给异地
	// EDU_UNABLE_TO_ATTEND_IN_PERSON (教育)无法线下参加
	// 若clue_convert_state =ARRIVAL，对于家居行业可选扩展标签，则该字段为非必填，可选值:
	// HOME_MEASUREMENT_COMPLETED [家居]已量房
	// HOME_ORDER_DISPATCHED [家居]已派单
	// 若clue_convert_state =CLUE_CONFIRM，可选扩展标签，则该字段为非必填，可选值:
	// 家居行业
	// HOME_APPOINTMENT_SCHEDULED [家居]已约时间
	// HOME_CAN_ADD_WECHAT [家居]可加微
	// HOME_CONNECTED_WITH_INTERESTED_PARTY [家居]接通有意向
	// 教育行业
	// EDU_FIRST_CLASS_ATTENDED [教育]首次到课
	// EDU_INTERESTED [教育]有意向
	// 若clue_convert_state =CONVERSION_CLASS，家居行业可选扩展标签，则该字段为非必填，可选值:
	// HOME_CONTRACT_SIGNED [家居]已签约
	// HOME_SERVICE_DELIVERY_COMPLETED [家居]已完成服务交付
	ReasonCode local.ClueReasonCode `json:"reason_code,omitempty"`
	// ReasonMessage 自定义原因信息，选填，当reason_code = OTHER时建议填写
	ReasonMessage string `json:"reason_message,omitempty"`
}

// Encode implements PostRequest interface
func (r LifeCallbackRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
