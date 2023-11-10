package platform

import (
	"encoding/json"
	"fmt"
	"github.com/aospace/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
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
	path := fmt.Sprintf("/platform/boxes/%v/route", c.BoxUUID)

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

	output := new(SpacePlatformMigrationOutResponse)
	err = utils.GetBody(resp, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
