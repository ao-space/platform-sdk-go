package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetBody(resp *http.Response, body interface{}) error {
	var respBody bytes.Buffer
	_, err := io.Copy(&respBody, resp.Body)
	if err != nil {
		return NewError(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		Err := new(Error)
		if err = json.Unmarshal([]byte(respBody.String()), Err); err != nil {
			return err
		}
		return Err
	}

	if err = json.Unmarshal([]byte(respBody.String()), body); err != nil {
		return NewError(err.Error())
	}

	return nil
}

func ToString(body interface{}) string {
	b, _ := json.Marshal(body)
	return string(b)
}
