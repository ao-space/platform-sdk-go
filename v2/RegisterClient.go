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
	if !c.IsAvailable(uriRegisterClient, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriRegisterClient)
	}
	uri := fmt.Sprintf("/platform/boxes/%v/users/%v/clients", c.BoxUUID, input.UserId)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPost, url)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)
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
