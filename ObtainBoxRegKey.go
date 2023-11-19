package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
	"time"
)

type ObtainBoxRegKeyRequest struct {
	BoxUUID    string   `json:"boxUUID"`
	ServiceIds []string `json:"serviceIds"`
	Sign       string   `json:"sign,optional"`
}

type ObtainBoxRegKeyResponse struct {
	BoxUUID      string         `json:"boxUUID"`
	TokenResults []TokenResults `json:"tokenResults"`
}

type TokenResults struct {
	ServiceId string    `json:"serviceId"`
	BoxRegKey string    `json:"boxRegKey"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// GetBoxRegKey 获取访问令牌
func (c *Client) ObtainBoxRegKey(input *ObtainBoxRegKeyRequest) (*ObtainBoxRegKeyResponse, error) {
	if !c.IsAvailable(uriObtainBoxRegKey, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriObtainBoxRegKey)
	}

	path := "/platform/auth/box_reg_keys"

	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodPost, URL)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)

	output := new(ObtainBoxRegKeyResponse)
	if err = utils.GetBody(resp, output); err != nil {
		return nil, err
	}

	c.BoxUUID = output.BoxUUID
	c.TokenResults = &output.TokenResults[0]

	return output, nil
}
