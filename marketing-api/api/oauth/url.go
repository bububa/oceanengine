package oauth

import (
	"fmt"
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/core"
)

// 生成授权链接
func Url(clt *core.SDKClient, redirectUrl string, state string) string {
	values := &url.Values{}
	values.Set("app_id", clt.AppID)
	if state != "" {
		values.Set("state", state)
	}
	values.Set("redirect_url", redirectUrl)
	return fmt.Sprintf("%s?%s", core.BASE_URL, values.Encode())
}
