package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// MaterialAttributesList 获取视频素材评估标签（新版）
// 本接口是「获取素材标签列表」、「获取素材标签信息」接口的升级版，支持查询账户下视频库的素材评估标签。新接口在支持原接口能力的基础上，扩展了以下新能力：
// - account_id支持4类平台账户：巨量广告/巨量千川/工作台（组织）/方舟，即可查询4类账户下的素材评估标签。如果您管理账户量级大、查询素材量级多，可使用本接口，by工作台/方舟账户拉取账户下素材的评估标签
//  1. 工作台账户维度（即管家账户/纵横组织账户，现已更名）：可获取查询「资产」-「视频库」下所有视频素材的评估标签
//     2.方舟账户维度：可获取方舟代理公司下的视频素材评估标签
//
// - 支持查询素材存在搬运风险标签：
//  1. 搬运授权保护的维度为广告账户，请使用巨量广告/巨量千川账户查询该标签。
//  2. 接口目前返回工作台/方舟账户维度下「存在搬运风险」标签，仅代表一个素材在工作台/方舟下的某个广告账户下存在搬运风险，暂不支持查询关联账户信息。
//
// 注意本接口支持在不同账户维度下查询素材的评估标签（创意生态标签介绍），能力支持情况如下：
// - 目前仅支持视频素材，可以获取的素材信息包括在投素材和不在投素材，查询结果不区分是否在投
// - 巨量广告或千川的「低效素材」、「同质化挤压严重」、「同质化排队素材」可搭配「创建素材清理任务」/「获取清理任务列表」/「下载清理任务结果」3个素材清理接口使用
func MaterialAttributesList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.MaterialAttributesListRequest) (*file.MaterialAttributesListResult, error) {
	var resp file.MaterialAttributesListResponse
	if err := clt.Get(ctx, "2/file/material_attributes/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
