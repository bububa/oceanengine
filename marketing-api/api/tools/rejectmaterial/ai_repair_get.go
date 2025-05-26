package rejectmaterial

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rejectmaterial"
)

// AIRepairGet 获取拒审素材修复建议
// 对应PC端产品能力：拒审素材「一键过审」产品手册一一过审真的「方便」
// 接口规则：
// 1. 默认返回最近7天、未采纳的修复建议，超过7天有效期的修复建议会失效（接口不会再返回）
// 2. 接口应答参数中拒审原视频素材的预览地址（preview_url）、修复后视频封面地址（ai_repair_video_cover_id）、修复后视频预览地址（ai_repair_preview_url）开启同主体校验，需要广告主授权时勾选敏感物料授权 / 开发者主体&广告主主体一致，方可正常获取
// 3. 本接口对应支持SPI事件（AD拒审素材修复成功事件），在系统完成拒审素材修复后，您可第一时间接收通知
func AIRepairGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *rejectmaterial.AIRepairGetRequest) (*rejectmaterial.AIRepairGetResult, error) {
	var resp rejectmaterial.AIRepairGetResponse
	if err := clt.GetAPI(ctx, "v3.0/reject_material/ai_repair/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
