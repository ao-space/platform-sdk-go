package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
)

// DeleteDevice 删除设备
func (c *Client) DeleteDevice() error {
	if !c.IsAvailable(uriDeleteDevice, http.MethodDelete) {
		return fmt.Errorf("the ability is not available: [%v] %v ", http.MethodDelete, uriDeleteDevice)
	}

	path := fmt.Sprintf("/platform/boxes/%v", c.BoxUUID)

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
