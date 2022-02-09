package track

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/track"
)

// Active API上报数据(new)
func Active(req *track.ActiveRequest) (string, error) {
	values := &url.Values{}
	if req.Callback != "" {
		values.Set("callback", req.Callback)
		values.Set("os", strconv.Itoa(req.Os))
		if req.Os == enum.Track_ANDROID {
			values.Set("imei", req.Imei)
			values.Set("muid", req.Muid)
			values.Set("oaid", req.Oaid)
			if req.OaidMd5 != "" {
				values.Set("oaid_md5", req.OaidMd5)
			}
		} else {
			values.Set("idfa", req.Idfa)
		}
		if req.Caid1 != "" {
			values.Set("caid1", req.Caid1)
		}
		if req.Caid2 != "" {
			values.Set("caid2", req.Caid2)
		}
	} else {
		values.Set("link", req.Link)
	}
	values.Set("event_type", strconv.Itoa(req.EventType))
	if req.ConvTime > 0 {
		values.Set("conv_time", strconv.FormatInt(req.ConvTime, 10))
	}
	if req.Source != "" {
		values.Set("source", req.Source)
	}
	for k, v := range req.Ext {
		values.Set(k, v)
	}
	var builder strings.Builder
	builder.WriteString("https://ad.oceanengine.com/track/activate/?")
	builder.WriteString(values.Encode())
	reqUrl := builder.String()
	resp, err := http.DefaultClient.Get(reqUrl)
	if err != nil {
		return reqUrl, err
	}
	defer resp.Body.Close()
	var ret track.Response
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return reqUrl, err
	}
	if ret.IsError() {
		return reqUrl, ret
	}
	return reqUrl, nil
}
