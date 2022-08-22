package grant_access

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type GrantAccess struct {
	gateway GrantAccessGateway
	grantAccessError error
}

func New(gateway GrantAccessGateway) GrantAccess {
	var grantAccessError error
	if gateway == nil {
		grantAccessError = internal_error.New("the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway")
	}
	return GrantAccess{
		gateway: gateway,
		grantAccessError: grantAccessError,
	}
}

func (usecase *GrantAccess) Grant(userId int64, resourceId int64, resource string, permissions []string) {
	accessControl := access_control.New(userId, resourceId, resource, permissions)
	usecase.grantAccessError = usecase.gateway.Grant(accessControl)
}

func (usecase GrantAccess) Error() error {
	return usecase.grantAccessError
}

func (usecase GrantAccess) HasError() bool {
	return usecase.grantAccessError != nil
}
