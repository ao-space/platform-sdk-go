package platform

import (
	"encoding/json"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type RegisterDeviceResponse struct {
	BoxUUID       string        `json:"boxUUID"`
	NetWorkClient netWorkClient `json:"networkClient"`
}

type netWorkClient struct {
	ClientId  string `json:"clientId"`
	SecretKey string `json:"secretKey"`
}

// RegisterDevice 注册设备
func (c *Client) RegisterDevice() (*RegisterDeviceResponse, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"boxUUID": c.BoxUUID,
	})
	if err != nil {
		return nil, err
	}
	c.SetOperation(http.MethodPost, c.BaseURL+"/platform/boxes")
	resp, err := c.Send(c.Operation, requestBody)
	output := RegisterDeviceResponse{}
	if err = utils.GetBody(resp, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
