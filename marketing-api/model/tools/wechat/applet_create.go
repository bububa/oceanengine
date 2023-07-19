package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AppletCreateRequest 创建微信小程序 API Request
type AppletCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 小程序名称, 须与小程序实际名称保持一致，最大长度不超过50
	Name string `json:"name,omitempty"`
	// UserName 小程序原始ID，获取方式：小程序后台>设置>基本设置>账号信息>原始ID，例如“gh_”，最大长度不超过20
	UserName string `json:"user_name,omitempty"`
	// Path 小程序路径参数，如有跳转指定页面或自定义参数监测广告效果数据诉求，需填写。获取方式：登录关联的公众号后台>素材管理>新建图文素材>小程序>选择小程序>获取更多路径
	Path string `json:"path,omitempty"`
	// RemarkMessage 资产备注信息，最大长度不超过15
	RemarkMessage string `json:"remark_message,omitempty"`
	// IconImageURL 小程序icon图片的url，尺寸60*60px,大小不超2MB
	IconImageURL string `json:"icon_image_url,omitempty"`
	// HeadImageURL 顶部头图的url，大小不超过5MB的图片，推荐尺寸2:1
	HeadImageURL string `json:"head_image_url,omitempty"`
	// Labels 小程序标签，每个标签长度不超过6，最多支持6个标签
	Labels []string `json:"labels,omitempty"`
	// GuideText 引导文案，最大长度不超过14
	GuideText string `json:"guide_text,omitempty"`
	// ImagesVerticalURL 小程序竖图的url地址，要求尺寸必须为3:5，否则报错。
	// 仅支持竖图和横图择其一类型上传，若同时上传两种类型的图片则报错
	// 小程序横图或竖图上传最少3张，最多8张，大小不超过5MB
	ImagesVerticalURL []string `json:"images_vertical_url,omitempty"`
	// ImagesHorizontalURL 小程序竖图的url地址，要求尺寸必须为3:5，否则报错。
	// 仅支持竖图和横图择其一类型上传，若同时上传两种类型的图片则报错
	// 小程序横图或竖图上传最少3张，最多8张，大小不超过5MB
	ImagesHorizontalURL []string `json:"images_horizontal_url,omitempty"`
	// Introduction 小程序简介，最大长度不超过50
	Introduction string `json:"introduction,omitempty"`
}

// Encode implements PostRequest interface
func (r AppletCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AppletCreateResponse 创建微信小程序 API Response
type AppletCreateResponse struct {
	model.BaseResponse
	Data struct {
		Data struct {
			// InstanceID 微信小程序资产ID
			InstanceID uint64 `json:"instance_id,omitempty"`
		} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
