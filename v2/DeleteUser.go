package platform

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
)

// DeleteUser 删除用户
func (c *Client) DeleteUser(userID string) error {
	if !c.IsAvailable(uriDeleteUser, http.MethodDelete) {
		return fmt.Errorf("the ability is not available: [%v] %v ", http.MethodDelete, uriDeleteUser)
	}

	path := fmt.Sprintf("/platform/boxes/%v/users/%v", c.BoxUUID, userID)

	URL := new(url.URL)
	copier.Copy(URL, c.BaseURL)
	URL = URL.JoinPath(path)

	op := new(Operation)
	op.SetOperation(http.MethodDelete, URL)

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
