package platform

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
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

	path := "/platform/ability"

	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodGet, URL)

	response, err := c.Send(op, nil)

	if err != nil {
		return nil, err
	}

	output := GetAbilityResponse{}
	err = utils.GetBody(response, &output)

	if err != nil {
		return nil, err
	}

	ability := make(map[string]map[string]map[int]int)

	for _, api := range output.PlatformApis {
		if ability[api.Uri] == nil {
			ability[api.Uri] = make(map[string]map[int]int)
		}
		if ability[api.Uri][strings.ToUpper(api.Method)] == nil {
			ability[api.Uri][strings.ToUpper(api.Method)] = make(map[int]int)
		}
		for _, version := range api.CompatibleVersions {
			ability[api.Uri][strings.ToUpper(api.Method)][version] = 1
		}
	}

	c.mu.Lock()
	c.Ability = ability
	c.mu.Unlock()

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

func (c *Client) IsAvailable(uri string, method string) bool {
	return c.Ability[uri][method][ApiVersion] == 1
}
