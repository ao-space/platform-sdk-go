package platform

import (
	"encoding/json"
	"fmt"
	"github.com/big-dust/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
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
	path := fmt.Sprintf("/platform/boxes/%v/users/%v/clients", c.BoxUUID, input.UserId)

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

	output := new(RegisterClientResponse)
	err = utils.GetBody(resp, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
