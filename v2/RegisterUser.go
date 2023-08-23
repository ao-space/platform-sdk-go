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
	if !c.IsAvailable(uriRegisterUser, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriRegisterUser)
	}
	uri := fmt.Sprintf("/platform/boxes/%v/users", c.BoxUUID)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPost, url)

	requestBody, _ := json.Marshal(input)
	response, err := c.Send(op, requestBody)
	if err != nil {
		return nil, err
	}
	output := RegisterUserResponse{}
	if err = utils.GetBody(response, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
