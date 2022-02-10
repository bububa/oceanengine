package oauth

import (
	"net/url"
	"strings"

	"github.com/bububa/oceanengine/marketing-api/core"
)

// Url 生成授权链接
func Url(clt *core.SDKClient, redirectUrl string, state string, materialAuth bool) string {
	values := &url.Values{}
	values.Set("app_id", clt.AppID)
	if state != "" {
		values.Set("state", state)
	}
	values.Set("redirect_uri", redirectUrl)
	if materialAuth {
		values.Set("material_auth", "1")
	}
	var builder strings.Builder
	builder.WriteString("https://ad.oceanengine.com/openapi/audit/oauth.html?")
	builder.WriteString(values.Encode())
	return builder.String()
}
