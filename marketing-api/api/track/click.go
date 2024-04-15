package track

import (
	"net/url"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// DEFAULT_CLICK_FIELDS 默认点击检测字段
var DEFAULT_CLICK_FIELDS = []string{
	"request_id",
	"aid",
	"promotion_id",
	"advertiser_id",
	"cid",
	"campaign_id",
	"project_id",
	"ctype",
	"csite",
	"imei",
	"idfa",
	"android_id",
	"oaid",
	"os",
	"caid",
	"mac",
	"mac1",
	"ip",
	"ua",
	"geo",
	"ts",
	"track_id",
	"callback",
	"callback_url",
	"model",
	"union_site",
}

// Click 生成击检测链接
func Click(baseUrl string, fields []string, adVersion model.AdVersion) string {
	if fields == nil {
		fields = DEFAULT_CLICK_FIELDS
	}
	parsedUrl, _ := url.Parse(baseUrl)
	values := parsedUrl.Query()
	for _, field := range fields {
		switch field {
		case "request_id":
			values.Set("request_id", "__REQUEST_ID__")
		case "aid":
			values.Set("aid", "__AID__")
		case "promotion_id":
			values.Set("promotion_id", "__PROMOTION_ID__")
		case "advertiser_id":
			values.Set("advertiser_id", "__ADVERTISER_ID__")
		case "cid":
			if adVersion != model.AdVersion_2 {
				values.Set("cid", "__CID__")
			}
		case "campaign_id":
			values.Set("campaign_id", "__CAMPAIGN_ID__")
		case "project_id":
			values.Set("project_id", "__PROJECT_ID__")
		case "ctype":
			values.Set("ctype", "__CTYPE__")
		case "csite":
			values.Set("csite", "__CSITE__")
		case "imei":
			values.Set("imei", "__IMEI__")
		case "idfa":
			values.Set("idfa", "__IDFA__")
		case "android_id":
			values.Set("android_id", "__ANDROIDID__")
		case "oaid":
			values.Set("oaid", "__OAID__")
		case "caid":
			values.Set("caid", "__CAID__")
		case "os":
			values.Set("os", "__OS__")
		case "mac":
			values.Set("mac", "__MAC__")
		case "mac1":
			values.Set("mac1", "__MAC1__")
		case "ip":
			values.Set("ip", "__IP__")
		case "ua":
			values.Set("ua", "__UA__")
		case "geo":
			values.Set("geo", "__GEO__")
		case "ts":
			values.Set("ts", "__TS__")
		case "track_id":
			values.Set("track_id", "__TRACK_ID__")
		case "callback":
			values.Set("callback", "__CALLBACK_PARAM__")
		case "callback_url":
			values.Set("callback_url", "__CALLBACK_URL__")
		case "model":
			values.Set("model", "__MODEL__")
		case "union_site":
			values.Set("union_site", "__UNION_SITE__")
		}
	}
	parsedUrl.RawQuery = values.Encode()
	return parsedUrl.String()
}
