package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBody(resp *http.Response, body interface{}) error {
	if resp == nil {
		return fmt.Errorf("response is nil")
	} else if resp.Body == nil {
		return fmt.Errorf("body is nil")
	}
	var respBody bytes.Buffer
	_, err := io.Copy(&respBody, resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		var Err Error
		if err = json.Unmarshal([]byte(respBody.String()), &Err); err != nil {
			return err
		}
		return &Err
	}
	if err = json.Unmarshal([]byte(respBody.String()), body); err != nil {
		return err
	}
	return nil
}

func ToString(body interface{}) string {
	b, _ := json.Marshal(body)
	return string(b)
}
