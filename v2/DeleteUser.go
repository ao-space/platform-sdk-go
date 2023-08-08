package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

// DeleteUser 删除用户
func (c *Client) DeleteUser(userID string) error {
	URL := c.BaseURL + fmt.Sprintf("/platform/boxes/%v/users/%v", c.BoxUUID, userID)
	c.SetOperation(http.MethodDelete, URL)
	resp, err := c.Send(c.Operation, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		Err := utils.GetBody(resp, nil)
		return Err
	}
	return nil
}
