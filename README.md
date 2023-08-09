# platform-sdk-go

## 简介

用于访问傲空间平台侧基础服务的Golang版SDK。

## 安装

1. 使用 go get 安装

   ```bash
   go get github.com/ao-space/platform-sdk-go/v2
   ```

2. 导入到你的代码

   ```bash
   import "github.com/ao-space/platform-sdk-go/v2"
   ```

## 快速开始

每个接口都有一个对应的 Request 结构和一个 Response 结构。例如获取访问令牌接口 ObtainBoxRegKey 有对应的请求结构体 ObtainBoxRegKeyRequest 和返回结构体 ObtainBoxRegKeyResponse 。

下面以获取访问令牌接口为例，介绍 SDK 的基础用法。

```go
package main

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
)

func main() {
    
	//创建客户端：需要指定平台侧基础服务的Host 和 选择是否手动设置transport
	client := platform.NewClientWithHost(platform.AoSpaceDomain, nil)
    
	//可选手动设置最近一次请求的requestId
	client.SetRequestId("XXXXX")
    
	//请求参数
	input := &platform.ObtainBoxRegKeyRequest{
		BoxUUID:    "XXXXX",
		ServiceIds: []string{"XXXXX"},
	}
	response, err := client.ObtainBoxRegKey(input)

	if err != nil {
		panic(err)
	}
	fmt.Println(utils.ToString(response))
}
```

## 文档

### SDK功能示例

1. 获取访问令牌

   - 用于空间平台认证设备身份，并生成访问令牌 box_reg_key，其它接口都需要在获取访问令牌后使用。

   ```go
   client = platform.NewClientWithHost("XXXXXX", nil)
   
   resp, err := client.ObtainBoxRegKey(&platform.ObtainBoxRegKeyRequest{
   	BoxUUID:    "XXXXX",
   	ServiceIds: []string{"XXXXX"},
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

2. 注册设备

   - 注册傲空间设备，空间平台为其分配 network client 信息。

   ```go
   resp, err := client.RegisterDevice()
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

3. 删除设备

   - 删除傲空间设备注册信息，包含用户注册信息、Client注册信息、网络资源等。

   ```go
   err := client.DeleteDevice()
   if err != nil {
   	fmt.Println(err)
   	return
   }
   ```

4. 注册用户

   - 注册用户，同步注册用户的绑定客户端

   ```go
   resp, err := client.RegisterUser(&platform.RegisterUserRequest{
   	UserID:     "XXX", //用户ID
   	Subdomain:  "XXX", //用户被指定的子域名
       UserType:   "XXX", //取值: user_admin,user_member
   	ClientUUID: "XXX", //客户端的 UUID
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

5. 申请用户域名

   - 申请用户的子域名，子域名全局唯一性

   ```go
   resp, err := client.GenerateUserDomain(&platform.GenerateUserDomainRequest{
   	EffectiveTime: "XXX", //有效期，单位秒，最长7天
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

6. 修改用户域名

   - 修改用户的子域名，仍然保留用户的历史域名

   ```go
   resp, err := client.ModifyUserDomain(&platform.ModifyUserDomainRequest{
   	UserId:    "XXX",
   	Subdomain: "XXX",
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

7. 删除用户

   - 删除用户注册信息，包含Client注册信息等

   ```go
   err := client.DeleteUser("your userId")
   if err != nil {
   	fmt.Println(err)
   	return
   }
   ```

8. 注册客户端

   - 注册客户端

   ```go
   resp, err := client.RegisterClient(&platform.RegisterClientRequest{
   	UserId:     "XXX",
   	ClientUUID: "XXX",
   	ClientType: "XXX", //客户端类型（绑定、扫码授权），取值：client_bind、client_auth
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

9. 删除客户端

   - 删除客户端注册信息

   ```go
   err := client.DeleteClient(&platform.DeleteClientRequest{
   	UserId:     "XXX",
   	ClientUUID: "XXX",
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   ```

10. 空间平台迁入

    - 用于向新空间平台迁入傲空间设备数据

    ```go
    resp, err := client.SpacePlatformMigration(&platform.SpacePlatformMigrationRequest{
    	NetworkClientId: "XXX",
    	UserInfos: []platform.UserMigrationInfo{
    		platform.UserMigrationInfo{
    			UserId:     "XXX",
    			UserDomain: "XXX",
    			UserType:   "XXX",
    			ClientInfos: []platform.ClientInfo{
    				platform.ClientInfo{
    					ClientUUID: "XXX",
    					ClientType: "XXX",
    				},
    			},
    		},
    	},
    })
    if err != nil {
    	fmt.Println(err)
    	return
    }
    ```

11. 空间平台迁出

    - 用于旧空间平台进行域名重定向

    ```go
    resp, err := client.SpacePlatformMigrationOut(&platform.SpacePlatformMigrationOutRequest{
    	UserDomainRouteInfos: []platform.UserDomainRouteInfo{
    		platform.UserDomainRouteInfo{
    			UserId:             "XXX",
    			UserDomainRedirect: "XXX", //重定向的用户域名
    		},
    	}})
    ```

### [API参考](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/API参考.md#API参考)

- [SDK初始化和配置](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/API参考.md#SDK初始化和配置)
- [主要操作的结构和方法](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/API参考.md#主要操作的结构和方法)
- [常量](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/API参考.md#常量)
- [错误类型](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/API参考.md#错误类型)
