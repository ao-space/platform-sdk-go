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
	if !c.IsAvailable(uriModifyUserDomainName, http.MethodPut) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPut, uriModifyUserDomainName)
	}
	uri := fmt.Sprintf("/platform/boxes/%v/users/%v/subdomain", c.BoxUUID, input.UserId)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPut, url)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)
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
