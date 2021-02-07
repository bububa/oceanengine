package core

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/bububa/oceanengine/marketing-api/core/internal/debug"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type SDKClient struct {
	AppID   string
	Secret  string
	debug   bool
	sandbox bool
}

func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		AppID:  appID,
		Secret: secret,
	}
}

func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

func (c *SDKClient) Post(gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqUrl := fmt.Sprintf("%s%s", BASE_URL, gw)
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

func (c *SDKClient) Get(gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	var reqQuery string
	if req != nil {
		reqQuery = req.Encode()
	}
	reqUrl := fmt.Sprintf("%s%s?%s", BASE_URL, gw, reqQuery)
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
