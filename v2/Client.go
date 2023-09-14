package platform

import (
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
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
	HTTPS     = "https"
)

type Client struct {
	HttpClient   *http.Client
	BoxUUID      string
	RequestId    string
	TokenResults *TokenResults
	BaseURL      *url.URL
	Ability      map[string]map[string]map[int]int
	mu           sync.Mutex
}

type Operation struct {
	Method string
	Url    string
}

// NewClientWithHost Host 是服务所在主机  transport 是连接池相关的配置
func NewClientWithHost(Host string, transport *http.Transport) (*Client, error) {

	if transport == nil {
		//默认连接池设置
		transport = NewDefaultTransport()
	}

	c := &Client{
		HttpClient: &http.Client{},
		mu:         sync.Mutex{},
	}

	c.SetBaseUrl(Host)

	_, err := c.GetAbility()

	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewDefaultTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:        5,                // 最大空闲连接数
		MaxIdleConnsPerHost: 2,                // 每个主机的最大空闲连接数
		IdleConnTimeout:     30 * time.Second, // 空闲连接超时时间
		TLSHandshakeTimeout: 10 * time.Second, // TLS握手超时时间
	}
}

func (c *Client) SetBaseUrl(Host string) {

	URL, _ := url.Parse(Host)
	if !URL.IsAbs() {
		URL.Scheme = HTTPS
		URL.Host = Host
	}

	URL.Path = "/v" + strconv.Itoa(ApiVersion)

	c.BaseURL = URL
}

func (c *Client) SetRequestId(requestId string) *Client {
	c.RequestId = requestId
	return c
}

func (c *Client) SetTransport(transport *http.Transport) {
	c.HttpClient.Transport = transport
}

func (op *Operation) SetOperation(method string, URL *url.URL) {
	op.Method = method
	op.Url = URL.String()
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

	request, _ := http.NewRequest(op.Method, op.Url, body)

	request.Header.Set("Accept", JSON)
	request.Header.Set("Content-Type", JSON)
	request.Header.Set("User-Agent", UserAgent)

	if c.RequestId == NULL {
		c.RequestId = uuid.New().String()
	}

	request.Header.Set("Request-Id", c.RequestId)

	if c.TokenResults.BoxRegKey != NULL {
		request.Header.Set("Box-Reg-key", c.TokenResults.BoxRegKey)
	}

	response, err := c.HttpClient.Do(request)

	if err != nil || response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		Logger.Error(time.Now().String()+": "+"request: ", request, " response: ", response)
	} else {
		Logger.Info(time.Now().String()+": "+"request: ", request, " response: ", response)
	}

	if err != nil {
		return nil, utils.NewError(err.Error())
	}

	c.RequestId = NULL
	return response, nil
}
