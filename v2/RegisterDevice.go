package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
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
	path := "/platform/boxes"

	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodPost, URL)

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
