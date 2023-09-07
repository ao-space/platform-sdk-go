package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

// DeleteDevice 删除设备
func (c *Client) DeleteDevice() error {
	if !c.IsAvailable(uriDeleteDevice, http.MethodDelete) {
		return fmt.Errorf("the ability is not available: [%v] %v ", http.MethodDelete, uriDeleteDevice)
	}

	uri := fmt.Sprintf("/platform/boxes/%v", c.BoxUUID)

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
