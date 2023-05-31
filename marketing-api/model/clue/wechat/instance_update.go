package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InstanceUpdateRequest 更新微信号码包 API Request
type InstanceUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 微信号码包ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// WechatList 要修改的微信号列表，最多不超过100个
	WechatList []string `json:"wechat_list,omitempty"`
	// Action 更新操作。允许值:
	// ENABLE: 启用
	// DISABLE: 暂停
	// DELETE: 删除
	// ADD: 添加
	Action string `json:"action,omitempty"`
}

// Encode implement PostRequest interface
func (r InstanceUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// InstanceUpdateResponse 更新微信号码包 API Response
type InstanceUpdateResponse struct {
	model.BaseResponse
	Data *InstanceUpdateResult `json:"data,omitempty"`
}

// InstanceUpdateResult
type InstanceUpdateResult struct {
	// SuccessWechatList 修改成功的微信号列表
	SuccessWechatList []string `json:"success_wechat_list,omitempty"`
	// FailWechatList 修改失败的微信号列表
	FailWechatList []struct {
		// Wechat 微信号码
		Wechat string `json:"wechat,omitempty"`
		// FailReason 失败原因
		FailReason string `json:"fail_reason,omitempty"`
	} `json:"fail_wechat_list,omitempty"`
}
