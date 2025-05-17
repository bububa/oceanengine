package sitetemplate

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/sitetemplate"
)

// Create 基于站点创建模板
// 可通过此接口基于已有落地页创建落地页模版
// 仅部分组件支持创建且编辑组件，使用不支持的组件创建模版会报错提示 “站点中存在不支持的组件类型”，具体关系如下：
// 支持创建模版且编辑的组件：视频组件、图片组件、组图组件、文本组件、富文本组件、表单组件、按钮组件（按钮、应用下载、智能电话、新版预约按钮）、直接发券组件
// 仅支持创建模版的组件：视频轮播组件、文字区块组件、图片区块组件、图形组件、倒计时组件、地图组件、裁剪路径组件、tab 组件组件、图标组件、程序化图片组件、程序化文本组件
// 落地页的每个发布版本都可以创建落地页模板，一个落地页站点可以创建多个落地页模板
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *sitetemplate.CreateRequest) (*sitetemplate.Template, error) {
	var resp sitetemplate.CreateResponse
	if err := clt.Post(ctx, "2/tools/site_template/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
