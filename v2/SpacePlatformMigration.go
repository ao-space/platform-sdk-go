package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type SpacePlatformMigrationRequest struct {
	NetworkClientId string              `json:"networkClientId"`
	UserInfos       []UserMigrationInfo `json:"userInfos"`
}

type UserMigrationInfo struct {
	UserId      string       `json:"userId"`
	UserDomain  string       `json:"userDomain"`
	UserType    string       `json:"userType"`
	ClientInfos []ClientInfo `json:"clientInfos"`
}

type ClientInfo struct {
	ClientUUID string `json:"clientUUID"`
	ClientType string `json:"clientType"`
}

type SpacePlatformMigrationResponse struct {
	BoxUUID       string              `json:"boxUUID"`
	NetworkClient netWorkClient       `json:"netWorkClient"`
	UserInfos     []UserMigrationInfo `json:"userInfos"`
}

func (c *Client) SpacePlatformMigration(input *SpacePlatformMigrationRequest) (*SpacePlatformMigrationResponse, error) {
	requestBody, _ := json.Marshal(input)
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/migration", c.BoxUUID)
	c.SetOperation(http.MethodPost, URL)
	resp, err := c.Send(c.Operation, requestBody)
	if err != nil {
		return nil, err
	}
	output := SpacePlatformMigrationResponse{}
	err = utils.GetBody(resp, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
