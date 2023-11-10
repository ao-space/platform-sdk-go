package platform

import (
	"encoding/json"
	"fmt"
	"github.com/aospace/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
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
	if !c.IsAvailable(uriSpacePlatformMigration, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriSpacePlatformMigration)
	}
	path := fmt.Sprintf("/platform/boxes/%v/migration", c.BoxUUID)

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

	output := new(SpacePlatformMigrationResponse)
	err = utils.GetBody(resp, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
