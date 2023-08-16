package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBody(resp *http.Response, body interface{}) error {
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
		return fmt.Errorf("%v: %v", Err.Code, Err.Message)
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
