package form

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/form"
)

// Update 更新表单
// 本接口用于表单的更新修改，主要可更新如下：
// - 基础信息模块
//   - 包含表单名称，不会在C端体现出来
//
// - 分层表单
//   - 分层内部的元素同样对应elements
//   - 结合字段：enable_layer + layer_submit_text
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
//   - 商家端额外通知用户的信息，如“24h内联系您“ 支持增量更新，目前不支持elements模块的修改。
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *form.UpdateRequest) (uint64, error) {
	var resp form.UpdateResponse
	if err := clt.Post(ctx, "2/clue/form/update/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.InstanceID, nil
}
