package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type GetStatusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (c *Client) GetStatus() (*GetStatusResponse, error) {
	if !c.IsAvailable(uriGetStatus, http.MethodGet) {
		return nil, fmt.Errorf("the ability is not available: [%v] %v ", http.MethodGet, uriGetStatus)
	}
	uri := "/platform/status"

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodGet, url)

	response, err := c.Send(op, nil)
	if err != nil {
		return nil, err
	}
	output := GetStatusResponse{}
	err = utils.GetBody(response, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
