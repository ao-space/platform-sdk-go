package platform

const (
	uriObtainBoxRegKey           = "/v2/platform/auth/box_reg_keys"
	uriRegisterDevice            = "/v2/platform/boxes"
	uriDeleteDevice              = "/v2/platform/boxes/{box_uuid}"
	uriRegisterUser              = "/v2/platform/boxes/{box_uuid}/users"
	uriGenerateUserDomainName    = "/v2/platform/boxes/{box_uuid}/subdomains"
	uriModifyUserDomainName      = "/v2/platform/boxes/{box_uuid}/users/{user_id}/subdomain"
	uriDeleteUser                = "/v2/platform/boxes/{box_uuid}/users/{user_id}"
	uriRegisterClient            = "/v2/platform/boxes/{box_uuid}/users/{user_id}/clients"
	uriDeleteClient              = "/v2/platform/boxes/{box_uuid}/users/{user_id}/clients/{client_uuid}"
	uriSpacePlatformMigration    = "/v2/platform/boxes/{box_uuid}/migration"
	uriSpacePlatformMigrationOut = "/v2/platform/boxes/{box_uuid}/route"
	uriGetStatus                 = "/v2/platform/status"
)
