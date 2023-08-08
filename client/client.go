package client

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	JSON      = "application/json"
	UserAgent = "go-sdk-v2"
	NULL      = ""
)

type Client struct {
	HttpClient *http.Client
	BoxUUID    string
	BoxRegKey  string
	BaseURL    string
	Operation  *Operation
}

type Operation struct {
	Method    string
	URL       string
	RequestId string
}

func (c *Client) SetOperation(method string, URL string) {
	c.Operation.Method = method
	c.Operation.URL = URL
}

func (c *Client) SetLogPath(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.Out = file
		return nil
	} else {
		Logger.Info("Failed to logs to file, using default stderr")
		return err
	}
}

// Send 通用请求发送
func (c *Client) Send(op *Operation, input []byte) (*http.Response, error) {
	var body io.Reader = nil
	if input != nil {
		body = strings.NewReader(string(input))
	}
	request, _ := http.NewRequest(op.Method, op.URL, body)
	request.Header.Set("Accept", JSON)
	request.Header.Set("Content-Type", JSON)
	request.Header.Set("User-Agent", UserAgent)

	if op.RequestId == NULL {
		//生成requestID
		op.RequestId = uuid.New().String()
	}
	request.Header.Set("Request-Id", op.RequestId)

	if c.BoxRegKey != NULL {
		request.Header.Set("Box-Reg-key", c.BoxRegKey)
	}
	response, err := c.HttpClient.Do(request)
	//打印日志
	if err != nil || response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		Logger.Error(time.Now().String()+": "+"request: ", request, " response: ", response)
	} else {
		Logger.Info(time.Now().String()+": "+"request: ", request, " response: ", response)
	}
	if err != nil {
		return nil, err
	}

	c.Operation.RequestId = NULL
	return response, nil
}
