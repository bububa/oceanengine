package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselAwemeGet 获取创编可用的抖音图文素材
// 该接口用于获取「巨量广告平台」创建广告时可用的抖音图文素材，您可以获取抖音图文素材的item_id，在广告创编接口使用。
// 您必须先和抖音号建立以下任意3种生效中的授权关系：抖音号授权、主页作品授权、单个作品授权，才可获取抖音图文素材
// 当抖音号下没有图文素材时，列表会返回为空
func CarouselAwemeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.CarouselAwemeGetRequest) (*file.CarouselAwemeGetResult, error) {
	var resp file.CarouselAwemeGetResponse
	if err := clt.GetAPI(ctx, "v3.0/file/carousel/aweme/get", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
