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
	if !c.IsAvailable(uriGenerateUserDomainName, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriGenerateUserDomainName)
	}
	uri := fmt.Sprintf("/platform/boxes/%v/subdomains", c.BoxUUID)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPost, url)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)
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
