package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type ModifyUserDomainRequest struct {
	UserId    string `json:"-"`
	Subdomain string `json:"subdomain"`
}

type ModifyUserDomainResponse struct {
	Success    bool     `json:"success"`
	BoxUUID    string   `json:"boxUUID,omitempty"`
	UserId     string   `json:"userId,omitempty"`
	Subdomain  string   `json:"subdomain,omitempty"`
	Code       int      `json:"code,omitempty"`
	Error      string   `json:"error,omitempty"`
	Recommends []string `json:"recommends,omitempty"`
}

func (c *Client) ModifyUserDomain(input *ModifyUserDomainRequest) (*ModifyUserDomainResponse, error) {
	requestBody, _ := json.Marshal(input)
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/users/%v/subdomain", c.BoxUUID, input.UserId)
	c.SetOperation(http.MethodPut, URL)
	resp, err := c.Send(c.Operation, requestBody)
	if err != nil {
		return nil, err
	}
	output := ModifyUserDomainResponse{}
	err = utils.GetBody(resp, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
