package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
	"time"
)

type ObtainBoxRegKeyRequest struct {
	BoxUUID    string   `json:"boxUUID"`
	ServiceIds []string `json:"serviceIds"`
	Sign       string   `json:"sign,optional"`
}

type ObtainBoxRegKeyResponse struct {
	BoxUUID      string         `json:"boxUUID"`
	TokenResults []tokenResults `json:"tokenResults"`
}

type tokenResults struct {
	ServiceId string    `json:"serviceId"`
	BoxRegKey string    `json:"boxRegKey"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// GetBoxRegKey 获取访问令牌
func (c *Client) ObtainBoxRegKey(input *ObtainBoxRegKeyRequest) (*ObtainBoxRegKeyResponse, error) {
	if !c.IsAvailable(uriObtainBoxRegKey, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriObtainBoxRegKey)
	}
	uri := "/platform/auth/box_reg_keys"

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPost, url)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)
	output := ObtainBoxRegKeyResponse{}
	if err = utils.GetBody(resp, &output); err != nil {
		return nil, err
	}
	c.BoxUUID = output.BoxUUID
	c.BoxRegKey = output.TokenResults[0].BoxRegKey
	return &output, nil
}
