package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoCoverSuggest 获取视频智能封面
// 通过此接口，用户可以获取针对素材视频推荐的智能封面。智能封面是通过提取视频关键帧筛选出推荐封面，帮助发现视频内优质封面素材。
// 推荐封面图片的数量是1-13个，对于相似度极高的封面图片会进行去重等处理，由实际的视频内容和时长决定。
// 智能封面不是实时获取，而需要先根据status判断封面获取的状态，然后再进行获取视频封面！
// 新上传素材存在同步延迟情况，建议等待2-3分钟再尝试操作获取视频智能封面！
// 获取封面素材仅用于当前广告主投放使用，不支持推送！
func VideoCoverSuggest(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoCoverSuggestRequest) (*file.VideoCoverSuggestResponseData, error) {
	var resp file.VideoCoverSuggestResponse
	err := clt.Get(ctx, "2/tools/video_cover/suggest", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
