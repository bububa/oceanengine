package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GameCreateRequest 创建微信小游戏 API Request
type GameCreateRequest struct {
	// AccountID 账户id，accout_type类型对应账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型
	// 允许值：BP 巨量纵横组织、 AD 广告主账号
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// CompanyID 公司主体id，通过【获取纵横组织下所有主体信息】接口获取，应用资质审核环节将以该主体作为拥有应用投放相关权利的主体进行审核，account_type=BP时必填
	CompanyID uint64 `json:"company_id,omitempty"`
	// UserName 小游戏原始ID，获取方式：小程序后台>设置>基本设置>账号信息>原始ID，例如“gh_”
	UserName string `json:"user_name,omitempty"`
	// Name 小游戏名称，最大长度不超过50
	Name string `json:"name,omitempty"`
	// Path 小游戏路径参数，如有跳转指定页面或自定义参数监测广告效果数据诉求，需填写。
	Path string `json:"path,omitempty"`
	// QuantificationID 资质id，account_type=AD时，资质id通过【获取投放资质信息】接口查询且必填；account_type=BP时，系统将自动审核投放资质，无需上传资质id
	// 注：account_type=BP时，请您确保您已经上传了投放资质，且您上传投放资质的广告主账户认证的公司主体，需与资产创建流程中选择的公司主体一致，以用于资质审核。若未提交投放资质，请通过【投放资质提交接口】提交资质，否则资产审核将不通过
	QuantificationID string `json:"quantification_id,omitempty"`
	// AnchorList 锚点组件信息
	AnchorList []Anchor `json:"anchor_list,omitempty"`
	// AntiAddictionURL 防沉迷提示URL，请上传游戏内包含限制游戏时间或建议合理安排游戏相关内容的截图。如防沉迷提示、防沉迷公共、健康系统相关内容。常见形式（包括不限于）：进入青少年模式，青少年防沉迷公告弹窗/提示，健康系统提示，强制休息提醒等
	AntiAddictionURL string `json:"anti_addiction_url,omitempty"`
	// ScreenRecordURL 游戏内容视频URL，请上传3分钟的游戏内实际画面操作录屏视频，要求MP4格式文件，最大不超过100M。横版16:9，分辨率不低于960540。竖版9:16，分辨率不低于540960
	ScreenRecordURL string `json:"screen_record_url,omitempty"`
	// RealNameURL 实名认证URL，请上传游戏内包含实名注册功能的截图链接
	RealNameURL string `json:"real_name_url,omitempty"`
	// AgeRemindURL 适龄提醒URL，请上传包含适龄提示图标的游戏界面截图，图标样式需与中国音像数字出版协会官方网站下载保持一致
	AgeRemindURL string `json:"age_remind_url,omitempty"`
}

// Encode implement PostRequest Interface
func (r GameCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// GameCreateResponse 创建微信小游戏 API Response
type GameCreateResponse struct {
	model.BaseResponse
	Data struct {
		Data *WechatGame `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
