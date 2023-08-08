package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

// DeleteDevice 删除设备
func (c *Client) DeleteDevice() error {
	path := fmt.Sprintf("/platform/boxes/%v", c.BoxUUID)
	c.SetOperation(http.MethodDelete, c.BaseURL+path)
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
