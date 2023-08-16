package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type RegisterUserRequest struct {
	UserID     string `json:"userId"`
	Subdomain  string `json:"subdomain"`
	UserType   string `json:"userType"`
	ClientUUID string `json:"clientUUID"`
}

type RegisterUserResponse struct {
	BoxUUID    string `json:"boxUUID"`
	UserID     string `json:"userId"`
	UserDomain string `json:"userDomain"`
	UserType   string `json:"userType"`
	ClientUUID string `json:"clientUUID"`
}

func (c *Client) RegisterUser(input *RegisterUserRequest) (*RegisterUserResponse, error) {
	requestBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/users", c.BoxUUID)
	c.SetOperation(http.MethodPost, URL)
	response, err := c.Send(c.Operation, requestBody)
	if err != nil {
		return nil, err
	}
	output := RegisterUserResponse{}
	if err = utils.GetBody(response, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
