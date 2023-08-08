package platform

import (
	"encoding/json"
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type SpacePlatformMigrationOutRequest struct {
	UserDomainRouteInfos []UserDomainRouteInfo `json:"userDomainRouteInfos"`
}

type UserDomainRouteInfo struct {
	UserId             string `json:"userId"`
	UserDomainRedirect string `json:"userDomainRedirect"`
}

type SpacePlatformMigrationOutResponse struct {
	BoxUUID              string                `json:"boxUUID"`
	UserDomainRouteInfos []UserDomainRouteInfo `json:"userDomainRouteInfos"`
}

func (c *Client) SpacePlatformMigrationOut(input *SpacePlatformMigrationOutRequest) (*SpacePlatformMigrationOutResponse, error) {
	requestBody, _ := json.Marshal(input)
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/route", c.BoxUUID)
	c.SetOperation(http.MethodPost, URL)
	resp, err := c.Send(c.Operation, requestBody)
	if err != nil {
		return nil, err
	}
	output := SpacePlatformMigrationOutResponse{}
	err = utils.GetBody(resp, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
