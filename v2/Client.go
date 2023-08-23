package platform

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
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
	RequestId  string
	BaseUrl    string
	Ability    *Ability
	mu         sync.Mutex
}

type Ability struct {
	PlatformApis []Api
	mu           sync.Mutex
}

type Operation struct {
	Method string
	Uri    string
}

// NewClientWithHost Host 是服务所在主机  transport 是连接池相关的配置
func NewClientWithHost(Host string, transport *http.Transport) (*Client, error) {

	if transport == nil {
		//默认连接池设置
		transport = &http.Transport{
			MaxIdleConns:        5,                // 最大空闲连接数
			MaxIdleConnsPerHost: 2,                // 每个主机的最大空闲连接数
			IdleConnTimeout:     30 * time.Second, // 空闲连接超时时间
			TLSHandshakeTimeout: 10 * time.Second, // TLS握手超时时间
		}
	}

	c := &Client{
		HttpClient: &http.Client{},
		Ability:    &Ability{},
		mu:         sync.Mutex{},
	}

	c.SetHost(Host)

	_, err := c.GetAbility()

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) SetHost(Host string) {
	c.BaseUrl = "https://" + Host + "/v" + strconv.Itoa(ApiVersion)
}

func (c *Client) SetRequestId(requestId string) *Client {
	c.RequestId = requestId
	return c
}

func (c *Client) SetTransport(transport *http.Transport) {
	c.HttpClient.Transport = transport
}

func (op *Operation) SetOperation(method string, Uri string) {
	op.Method = method
	op.Uri = Uri
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

	request, _ := http.NewRequest(op.Method, op.Uri, body)

	request.Header.Set("Accept", JSON)
	request.Header.Set("Content-Type", JSON)
	request.Header.Set("User-Agent", UserAgent)

	if c.RequestId == NULL {
		c.RequestId = uuid.New().String()
	}

	request.Header.Set("Request-Id", c.RequestId)

	if c.BoxRegKey != NULL {
		request.Header.Set("Box-Reg-key", c.BoxRegKey)
	}

	c.mu.Lock()
	response, err := c.HttpClient.Do(request)
	c.mu.Unlock()

	if err != nil || response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		Logger.Error(time.Now().String()+": "+"request: ", request, " response: ", response)
	} else {
		Logger.Info(time.Now().String()+": "+"request: ", request, " response: ", response)
	}

	if err != nil {
		return nil, err
	}

	c.RequestId = NULL
	return response, nil
}
