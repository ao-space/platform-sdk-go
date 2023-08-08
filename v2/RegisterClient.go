package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type RegisterClientRequest struct {
	UserId     string `json:"-"`
	ClientUUID string `json:"clientUUID"`
	ClientType string `json:"clientType"`
}

type RegisterClientResponse struct {
	BoxUUID    string `json:"boxUUID"`
	UserId     string `json:"userId"`
	ClientUUID string `json:"clientUUID"`
	ClientType string `json:"clientType"`
}

func (c *Client) RegisterClient(input *RegisterClientRequest) (*RegisterClientResponse, error) {
	requestBody, _ := json.Marshal(input)
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/users/%v/clients", c.BoxUUID, input.UserId)
	c.SetOperation(http.MethodPost, URL)
	resp, err := c.Send(c.Operation, requestBody)
	if err != nil {
		return nil, err
	}
	output := RegisterClientResponse{}
	err = utils.GetBody(resp, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
