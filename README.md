English | [简体中文](https://github.com/big-dust/platform-sdk-go/blob/dev/README_cn.md)

# platform-sdk-go

## Introduction

Golang version SDK for accessing the fundamental services on the Ao.Space platform.

## Installation

1. Install using `go get`

   ```bash
   go get github.com/ao-space/platform-sdk-go/v2
   ```

2. Import into your code

   ```bash
   import "github.com/ao-space/platform-sdk-go/v2"
   ```

## Quick Start

Each interface has a corresponding Request structure and a Response structure. For example, the `ObtainBoxRegKey` interface has corresponding request and response structures named `ObtainBoxRegKeyRequest` and `ObtainBoxRegKeyResponse`, respectively.

Below is an example of how to use the SDK for obtaining an access token.

```go
package main

import (
	"fmt"
	"github.com/ao-space/platform-sdk-go/utils"
	"github.com/ao-space/platform-sdk-go/v2"
)

func main() {
    
	// Create a client: specify the Host of the platform's fundamental
    // service and optionally set the transport manually
	client := platform.NewClientWithHost(platform.AoSpaceDomain, nil)
    
	// Optionally set the request ID for the most recent request
	client.SetRequestId("XXXXX")
    
	// Request parameters
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

## Documentation

### SDK Functionality Examples

1. Obtain Box_Reg_Key

   - Used to authenticate the identity of the device on the space platform and generate `box_reg_keys`.

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

2. Register Device

   - Register AO.space device, and the space platform assigns network client information to it

   ```go
   resp, err := client.RegisterDevice()
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

3. Delete Device

   - Delete the registration information of AO.space device, including user registration information, client registration information, network resources, etc

   ```go
   err := client.DeleteDevice()
   if err != nil {
   	fmt.Println(err)
   	return
   }
   ```

4. Register User

   - Register users and synchronize their binding clients

   ```go
   resp, err := client.RegisterUser(&platform.RegisterUserRequest{
   	UserID:     "XXX", // User ID
   	Subdomain:  "XXX", //The subdomain name specified by the user
       UserType:   "XXX", //User type (administrator, member), value: user_admin、user_member
   	ClientUUID: "XXX", //The UUID of the client
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

5. Generate User Domain Name

   - Generate the user's subdomain name, and the subdomain name is unique globally

   ```go
   resp, err := client.GenerateUserDomain(&platform.GenerateUserDomainRequest{
   	EffectiveTime: "XXX", //Validity period, in seconds, up to 7 days
   })
   if err != nil {
   	fmt.Println(err)
   	return
   }
   fmt.Println(resp)
   ```

6. Modify User Domain Name

   - Modify the user's subdomain name, still retaining the user's historical subdomain name

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

7. Delete User

   - Delete user registration information, including client registration information, etc

   ```go
   err := client.DeleteUser("your userId")
   if err != nil {
   	fmt.Println(err)
   	return
   }
   ```

8. Register Client

   - Register Client

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

9. Delete Client

   - Delete client registration information

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

10. Space Platform Migration

    - Used to migrate AO.space device data to the new space platform

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

11. Space Platform Migration Out

    - Used for domain name redirection on old space platforms

    ```go
    resp, err := client.SpacePlatformMigrationOut(&platform.SpacePlatformMigrationOutRequest{
    	UserDomainRouteInfos: []platform.UserDomainRouteInfo{
    		platform.UserDomainRouteInfo{
    			UserId:             "XXX",
    			UserDomainRedirect: "XXX", //Redirected user domain name
    		},
    	}})
    ```

### API Reference

- [SDK Initialization and Configuration](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/en/API%20Reference.md#1-sdk-initialization-and-configuration)
- [Main Operation Structures and Methods](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/en/API%20Reference.md#2-main-operation-structures-and-methods)
- [Constants](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/en/API%20Reference.md#3-constants)
- [Error Code](https://github.com/big-dust/platform-sdk-go/blob/dev/docs/en/API%20Reference.md#4-error-code)

