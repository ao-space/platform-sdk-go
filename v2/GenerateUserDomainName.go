package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type GenerateUserDomainRequest struct {
	EffectiveTime string `json:"effectiveTime"`
}

type GenerateUserDomainResponse struct {
	BoxUUID   string `json:"boxUUID"`
	Subdomain string `json:"subdomain"`
	ExpiresAt string `json:"expiresAt"`
}

func (c *Client) GenerateUserDomain(input *GenerateUserDomainRequest) (*GenerateUserDomainResponse, error) {
	requestBody, _ := json.Marshal(input)
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/subdomains", c.BoxUUID)
	c.SetOperation(http.MethodPost, URL)
	resp, err := c.Send(c.Operation, requestBody)
	if err != nil {
		return nil, err
	}
	output := GenerateUserDomainResponse{}
	err = utils.GetBody(resp, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
