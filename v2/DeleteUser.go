package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"net/http"
)

// DeleteUser 删除用户
func (c *Client) DeleteUser(userID string) error {
	if !c.IsAvailable(uriDeleteUser, http.MethodDelete) {
		return fmt.Errorf("the ability is not available: [%v] %v ", http.MethodDelete, uriDeleteUser)
	}

	uri := fmt.Sprintf("/platform/boxes/%v/users/%v", c.BoxUUID, userID)

	url := c.BaseUrl + uri
	op := new(Operation)
	op.SetOperation(http.MethodDelete, url)

	resp, err := c.Send(op, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		Err := utils.GetBody(resp, nil)
		return Err
	}
	return nil
}
