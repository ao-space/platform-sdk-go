package platform

import (
	"fmt"
	"github.com/big-dust/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
)

type DeleteClientRequest struct {
	UserId     string
	ClientUUID string
}

func (c *Client) DeleteClient(input *DeleteClientRequest) error {
	if !c.IsAvailable(uriDeleteClient, http.MethodDelete) {
		return fmt.Errorf("the ability is not available: [%v] %v ", http.MethodDelete, uriDeleteClient)
	}

	path := fmt.Sprintf("platform/boxes/%v/users/%v/clients/%v", c.BoxUUID, input.UserId, input.ClientUUID)
	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodDelete, URL)
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
