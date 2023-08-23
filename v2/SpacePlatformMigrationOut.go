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
	if !c.IsAvailable(uriSpacePlatformMigrationOut, http.MethodPost) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodPost, uriSpacePlatformMigrationOut)
	}
	uri := fmt.Sprintf("/platform/boxes/%v/route", c.BoxUUID)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodPost, url)

	requestBody, _ := json.Marshal(input)
	resp, err := c.Send(op, requestBody)
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
