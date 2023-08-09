# API参考

## 1. SDK初始化和配置

- **NewClientWithHost(Host string, transport \*http.Transport) \*Client**：创建一个用于调用接口的client实例。

  **Host**: 平台侧服务主机域名。

  **transport**: http的连接池配置，默认参数为：

  ```go
  transport := &http.Transport{
  	MaxIdleConns:        5,               
  	MaxIdleConnsPerHost: 2,               
  	IdleConnTimeout:     30 * time.Second, 
  	TLSHandshakeTimeout: 10 * time.Second, 
  }
  ```

- **Client** : 集中平台侧基础服务接口调用方法。
- **(c *Client) SetHost(Host string)**：设置平台侧服务的主机

- **(c \*Client) SetLogPath(path string) error** : 设置日志输出地址 path ，并返回一个错误信息，默认为: ./sdk.log 。
- **(c \*Client) SetRequestId(requestId string) \*Client** : 手动设置当前一个请求的requestId，默认自动生成。
- **(c \*Client) SetTransport(transport \*http.Transport)** : 初始化后设置http连接池配置。

## 2. 主要操作的结构和方法

#### 2.1 获取访问令牌

- **(c \*Client) ObtainBoxRegKey(input \*ObtainBoxRegKeyRequest) ( \*ObtainBoxRegKeyResponse, error)**

  获取访问令牌的功能方法。

- **ObtainBoxRegKeyRequest**

  ```go
  type ObtainBoxRegKeyRequest struct {
  	BoxUUID    string   `json:"boxUUID"` 	   //设备的 UUID
  	ServiceIds []string `json:"serviceIds"`    //平台id：空间平台（serviceId=10001）
  	Sign       string   `json:"sign,optional"` //签名，使用公钥验证设备身份时必传
  }
  ```

- **ObtainBoxRegKeyResponse**

  ```go
  type ObtainBoxRegKeyResponse struct {
  	BoxUUID      string         `json:"boxUUID"`      //设备的 UUID
  	TokenResults []tokenResults `json:"tokenResults"` //设备的访问令牌
  }
  ```

- **tokenResults**

  ```go
  type tokenResults struct {
  	ServiceId string    `json:"serviceId"` //平台id
  	BoxRegKey string    `json:"boxRegKey"` //设备的访问令牌
  	ExpiresAt time.Time `json:"expiresAt"` //令牌有效时间
  }
  ```

#### 2.2 注册设备

- **(c \*Client) RegisterDevice() (\*RegisterDeviceResponse, error)**

  注册设备的功能方法。

- **RegisterDeviceResponse**

  ```go
  type RegisterDeviceResponse struct {
  	BoxUUID       string        `json:"boxUUID"`       //设备的 UUID
  	NetWorkClient netWorkClient `json:"networkClient"` //为其分配 network client 信息
  }
  ```

- **netWorkClient**

  ```go
  type netWorkClient struct {
  	ClientId  string `json:"clientId"`  //network 的客户端 ID
  	SecretKey string `json:"secretKey"` //访问密钥
  }
  ```

#### 2.3 删除设备

- **(c \*Client) DeleteDevice() error**

  删除设备的功能方法。

#### 2.4 注册用户

- **(c \*Client) RegisterUser(input \*RegisterUserRequest) (\*RegisterUserResponse, error)**

  注册用户的功能方法。

- **RegisterUserRequest**

  ```go
  type RegisterUserRequest struct {
  	UserID     string `json:"userId"`     //用户的 ID
  	Subdomain  string `json:"subdomain"`  //用户被指定的子域名
  	UserType   string `json:"userType"`   //用户类型（管理员、普通成员），取值：user_admin、user_member
  	ClientUUID string `json:"clientUUID"` //客户端的 UUID
  }
  ```

- **RegisterUserResponse**

  ```go
  type RegisterUserResponse struct {
  	BoxUUID    string `json:"boxUUID"`    //设备的 UUID
  	UserID     string `json:"userId"`     //用户的 ID
  	UserDomain string `json:"userDomain"` //为用户分配的用户域名，该域名可以用于后续的业务访问
  	UserType   string `json:"userType"`   //用户类型（管理员、普通成员）
  	ClientUUID string `json:"clientUUID"` //客户端的 UUID
  }
  ```

#### 2.5 申请用户域名

- **(c \*Client) GenerateUserDomain(input \*GenerateUserDomainRequest) (\*GenerateUserDomainResponse, error)**

  申请用户域名的功能方法。

- **GenerateUserDomainRequest**

  ```go
  type GenerateUserDomainRequest struct {
  	EffectiveTime string `json:"effectiveTime"` //有效期，单位秒，最长7天
  }
  ```

- **GenerateUserDomainResponse**

  ```go
  type GenerateUserDomainResponse struct {
  	BoxUUID   string `json:"boxUUID"`   //设备的 UUID
  	Subdomain string `json:"subdomain"` //用户被指定的子域名
  	ExpiresAt string `json:"expiresAt"` //有效期
  }
  ```

#### 2.6 修改用户域名

- **(c \*Client) ModifyUserDomain(input \*ModifyUserDomainRequest) (\*ModifyUserDomainResponse, error)**

  修改用户域名的功能方法。

- **ModifyUserDomainRequest**

  ```go
  type ModifyUserDomainRequest struct {
  	UserId    string `json:"-"`         //用户的 ID
  	Subdomain string `json:"subdomain"` //用户指定的新的子域名
  }
  ```

- **ModifyUserDomainResponse**

  ```go
  type ModifyUserDomainResponse struct {
     Success    bool     `json:"success"`			     //是否成功
     BoxUUID    string   `json:"boxUUID,omitempty"`    //设备的 UUID, success 为 true 时返回
     UserId     string   `json:"userId,omitempty"`     //用户的 ID, success 为 true 时返回
     Subdomain  string   `json:"subdomain,omitempty"`  //用户指定的新的子域名, success 为 true 时返回
     Code       int      `json:"code,omitempty"`       //错误码, success 为 false 时返回
     Error      string   `json:"error,omitempty"`      //错误消息, success 为 false 时返回
     Recommends []string `json:"recommends,omitempty"` //推荐的subdomain, success 为 false 时返回
  }
  ```

#### 2.7 删除用户

- **DeleteUser(userId string) error** 

  删除指定用户的功能方法，需要接收一个 userID 用户id 参数。

#### 2.8 注册客户端

- **(c \*Client) RegisterClient(input \*RegisterClientRequest) (\*RegisterClientResponse, error)**

  注册客户端的功能方法。

- **RegisterClientRequest**

  ```go
  type RegisterClientRequest struct {
  	UserId     string `json:"-"`		  //用户的 ID
  	ClientUUID string `json:"clientUUID"` //客户端的 UUID
  	ClientType string `json:"clientType"` //客户端类型（绑定、扫码授权），取值：client_bind、client_auth
  }
  ```

- **RegisterClientResponse**

  ```go
  type RegisterClientResponse struct {
  	BoxUUID    string `json:"boxUUID"`    //设备的 UUID
  	UserId     string `json:"userId"`     //用户的 ID
  	ClientUUID string `json:"clientUUID"` //客户端的 UUID
  	ClientType string `json:"clientType"` //客户端类型（绑定、扫码授权
  }
  ```

#### 2.9 删除客户端

- **(c \*Client) DeleteClient(input \*DeleteClientRequest) error**

  删除客户端的功能方法。

- **DeleteClientRequest**

  ```go
  type DeleteClientRequest struct {
  	UserId     string //用户的 ID
  	ClientUUID string //客户端的 UUID
  }
  ```

#### 2.10 空间平台迁入

- **(c *Client) SpacePlatformMigration(input \*SpacePlatformMigrationRequest) (\*SpacePlatformMigrationResponse, error)**

  向空间平台迁入的功能方法。

- **SpacePlatformMigrationRequest**

  ```go
  type SpacePlatformMigrationRequest struct {
  	NetworkClientId string              `json:"networkClientId"` //network 的客户端 ID
  	UserInfos       []UserMigrationInfo `json:"userInfos"`       //用户列表
  }
  ```

- **UserMigrationInfo**

  ```go
  type UserMigrationInfo struct {
     UserId      string       `json:"userId"`      // 用户的 ID
     UserDomain  string       `json:"userDomain"`  //用户域名
     UserType    string       `json:"userType"`    //用户类型（管理员、普通成员）
     ClientInfos []ClientInfo `json:"clientInfos"` //Client 列表
  }
  ```

- **ClientInfo**

  ```go
  type ClientInfo struct {
     ClientUUID string `json:"clientUUID"` //客户端的 UUID
     ClientType string `json:"clientType"` //客户端类型（绑定、扫码授权），取值：client_bind、client_auth
  }
  ```

- **SpacePlatformMigrationResponsed**

  ```go
  type SpacePlatformMigrationResponse struct {
     BoxUUID       string              `json:"boxUUID"`       //设备的 UUID
     NetworkClient netWorkClient       `json:"netWorkClient"` //为设备分配的 network client 信息
     UserInfos     []UserMigrationInfo `json:"userInfos"`     //用户列表
  }
  ```

#### 2.11 空间平台迁出

- **(c \*Client) SpacePlatformMigrationOut(input \*SpacePlatformMigrationOutRequest) (*SpacePlatformMigrationOutResponse, error)** 

  迁出空间平台的功能方法。

- **SpacePlatformMigrationOutRequest**

  ```go
  type SpacePlatformMigrationOutRequest struct {
     UserDomainRouteInfos []UserDomainRouteInfo `json:"userDomainRouteInfos"` //用户域名映射关系
  }
  ```

- **UserDomainRouteInfo**

  ```go
  type UserDomainRouteInfo struct {
     UserId             string `json:"userId"`             //用户的 ID
     UserDomainRedirect string `json:"userDomainRedirect"` //重定向的用户域名
  }
  ```

- **SpacePlatformMigrationOutResponse**

  ```go
  type SpacePlatformMigrationOutResponse struct {
     BoxUUID              string                `json:"boxUUID"`              //设备的 UUID
     UserDomainRouteInfos []UserDomainRouteInfo `json:"userDomainRouteInfos"` //用户域名映射关系
  }
  ```

## 3. 常量

- **ApiVersion**：api版本
- **AoSpaceDomain**：傲空间域名

## 4. 错误类型

| 错误码   | 错误信息                           | 说明                  |
| -------- | ---------------------------------- | --------------------- |
| SSP-2012 | input parameter:{0} error          | 请求参数错误          |
| SSP-2017 | subdomain does not exist           | 子域名不存在          |
| SSP-2018 | subdomain already exist            | 子域名已存在          |
| SSP-2019 | subdomain already used             | 子域名已使用          |
| SSP-2020 | reach subdomain upper limit        | 子域名数量已达到上限  |
| SSP-2021 | box uuid has already registered    | box uuid 已注册       |
| SSP-2022 | box uuid had not registered        | box uuid 未注册       |
| SSP-2023 | user id has already registered     | user id 已注册        |
| SSP-2024 | user id has not registered         | user id 未注册        |
| SSP-2025 | client uuid has already registered | client uuid 已注册    |
| SSP-2026 | client uuid has not registered     | client uuid 未注册    |
| SSP-2028 | network client does not exist      | network client 不存在 |
| SSP-2049 | network server does not exist      | network server 不存在 |
| SSP-2050 | subdomain is not in use            | 子域名未使用          |
| SSP-2051 | subdomain is reserved              | 子域名不合法          |
| SSP-2060 | migration in acquire lock error    | 迁入操作获取锁失败    |
| SSP-2061 | migration out acquire lock error   | 迁出操作获取锁失败    |