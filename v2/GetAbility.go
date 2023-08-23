package platform

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
	"strings"
	"time"
)

type Api struct {
	Method             string `json:"method"`
	Uri                string `json:"uri"`
	BriefUri           string `json:"briefUri"`
	CompatibleVersions []int  `json:"compatibleVersions"`
	Type               string `json:"type"`
	Desc               string `json:"desc"`
}

type GetAbilityResponse struct {
	PlatformApis []Api `json:"platformApis"`
}

func (c *Client) GetAbility() (*GetAbilityResponse, error) {

	uri := "/platform/ability"

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodGet, url)

	response, err := c.Send(op, nil)

	if err != nil {
		return nil, err
	}

	output := GetAbilityResponse{}
	err = utils.GetBody(response, &output)

	if err != nil {
		return nil, err
	}

	c.Ability.mu.Lock()
	c.Ability.PlatformApis = output.PlatformApis
	c.Ability.mu.Unlock()

	return &output, nil
}

func (c *Client) FlushAbilityWithDuration(duration time.Duration) func() {
	return func() {
		for {
			_, err := c.GetAbility()
			if err != nil {
				break
			}
			time.Sleep(duration)
		}
	}
}

func (c *Client) IsAvailable(Uri string, method string) bool {
	for _, api := range c.Ability.PlatformApis {
		if api.Uri != Uri {
			continue
		}
		if api.Method != strings.ToLower(method) {
			continue
		}
		for _, compatibleVersion := range api.CompatibleVersions {
			if compatibleVersion == ApiVersion {
				return true
			}
		}
	}
	return false
}
