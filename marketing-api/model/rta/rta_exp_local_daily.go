package rta

import (
	"encoding/json"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
	"strconv"
)

type GetRtaExpLocalDailyReq struct {
	AdvertiserID uint64  `json:"advertiser_id"`
	RtaID        int     `json:"rta_id"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	VID          int     `json:"vid,omitempty"`
	CusVID       int     `json:"cus_vid,omitempty"`
	Filtering    *Filter `json:"filtering,omitempty"`
}

type Filter struct {
	BidCoef []string `json:"bid_coef"`
}

func (r GetRtaExpLocalDailyReq) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.RtaID > 0 {
		values.Set("rta_id", strconv.Itoa(r.RtaID))
	}
	if r.StartDate != "" {
		values.Set("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		values.Set("end_date", r.EndDate)
	}

	if r.VID > 0 {
		values.Set("vid", strconv.Itoa(r.VID))
	}
	if r.CusVID > 0 {
		values.Set("cus_vid", strconv.Itoa(r.CusVID))
	}
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

type GetRtaExpLocalDailyResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		Data []*GetRtaExpLocalDailyData `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
type GetRtaExpLocalDailyData struct {
	Click   int     `json:"click"`
	Convert int     `json:"convert"`
	Cost    float64 `json:"cost"`
	CusVid  int     `json:"cus_vid"`
	Date    string  `json:"date"`
	Show    int     `json:"show"`
	Vid     int     `json:"vid"`
}
