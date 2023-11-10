package platform

import (
	"encoding/json"
	"fmt"
	"github.com/aospace/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
)

type GenerateUserDomainRequest struct {
	EffectiveTime int `json:"effectiveTime"`
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

	path := fmt.Sprintf("platform/boxes/%v/subdomains", c.BoxUUID)
	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodPost, URL)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)
	if err != nil {
		return nil, err
	}

	output := new(GenerateUserDomainResponse)
	err = utils.GetBody(resp, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
