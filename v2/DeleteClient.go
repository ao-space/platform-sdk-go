package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

type DeleteClientRequest struct {
	UserId     string
	ClientUUID string
}

func (c *Client) DeleteClient(input *DeleteClientRequest) error {
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/users/%v/clients/%v", c.BoxUUID, input.UserId, input.ClientUUID)
	c.SetOperation(http.MethodDelete, URL)
	resp, err := c.Send(c.Operation, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		err = utils.GetBody(resp, nil)
		return err
	}
	return nil
}
