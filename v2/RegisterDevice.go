package platform

import (
	"encoding/json"
	"fmt"
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
	if !c.IsAvailable(uriRegisterDevice, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriRegisterDevice)
	}
	uri := "/platform/boxes"

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPost, url)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"boxUUID": c.BoxUUID,
	})
	resp, err := c.Send(op, requestBody)
	if err != nil {
		return nil, err
	}
	output := RegisterDeviceResponse{}
	if err = utils.GetBody(resp, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
