package platform

import (
	"github.com/ao-space/platform-sdk-go/client"
	"net/http"
	"time"
)

type Client struct {
	*client.Client
}

// NewClientWithHost Host 是服务所在主机  transport 是连接池相关的配置
func NewClientWithHost(Host string, transport *http.Transport) *Client {
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
		&client.Client{
			HttpClient: &http.Client{
				Transport: transport,
			},
			Operation: &client.Operation{},
		},
	}
	c.SetHost(Host)
	return c
}

func (c *Client) SetHost(Host string) {
	c.BaseURL = "https://" + Host + "/" + ApiVersion
}

func (c *Client) SetRequestId(requestId string) *Client {
	c.Operation.RequestId = requestId
	return c
}

func (c *Client) SetTransport(transport *http.Transport) {
	c.HttpClient.Transport = transport
}
