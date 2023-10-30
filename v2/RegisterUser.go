package platform

import (
	"encoding/json"
	"fmt"
	"github.com/big-dust/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
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
	path := fmt.Sprintf("/platform/boxes/%v/users", c.BoxUUID)

	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodPost, URL)

	requestBody, _ := json.Marshal(input)
	response, err := c.Send(op, requestBody)
	if err != nil {
		return nil, err
	}

	output := new(RegisterUserResponse)
	if err = utils.GetBody(response, output); err != nil {
		return nil, err
	}

	return output, nil
}
