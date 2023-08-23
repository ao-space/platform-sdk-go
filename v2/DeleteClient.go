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
	if !c.IsAvailable(uriDeleteClient, http.MethodDelete) {
		return fmt.Errorf("the ability is not available: [%v] %v ", http.MethodDelete, uriDeleteClient)
	}
	uri := fmt.Sprintf("/platform/boxes/%v/users/%v/clients/%v", c.BoxUUID, input.UserId, input.ClientUUID)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodDelete, url)
	resp, err := c.Send(op, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = utils.GetBody(resp, nil)
		return err
	}
	return nil
}
