package form

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/form"
)

// Create 创建表单
// 本接口用于创建表单，主要包括以下信息的设置：
// - 基础信息模块
//   - 包含表单名称，不会在C端体现出来
//
// - 表单元素配置
//   - 对应FormElement
//   - 支持添加多种类型元素：姓名、电话、邮箱、数值、性别、日期、城市、自定义文本、多文本
//
// - 分层表单
//   - 分层内部的元素同样对应elements
//   - 结合字段：enable_layer + layer_submit_text
//
// - 验证类型配置
//   - validate_type字段，主要指C端是否进行短信验证
//   - 当表单含有电话类型字段才生效
//
// - 高级设置
//   - 高级设置属于扩展功能，目前主要包含三项内容：
//   - 计数设置
//   - 包含递增和递减两种类型
//   - 基于计数前文案+计数值+计数后文案拼接完整方案
//   - 展示报名用户：
//   - 包含滚动墙和滚动条两种类型
//   - 基于已提交用户信息展示数据
//   - 提交成功提示语：
//   - 商家端额外通知用户的信息，如“24h内联系您“
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *form.CreateRequest) (uint64, error) {
	var resp form.CreateResponse
	if err := clt.Post(ctx, "2/clue/form/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.InstanceID, nil
}
